package repository

import (
	"context"
	"log/slog"
	"my-wallet-ntier-mongo/constant"
	"my-wallet-ntier-mongo/model"
	transactionResponse "my-wallet-ntier-mongo/response"
	"my-wallet-ntier-mongo/utils"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type TransactionRepository struct {
	mongoDB *mongo.Database
}

func NewTransactionRepository(mongoDB *mongo.Database) *TransactionRepository {
	return &TransactionRepository{mongoDB: mongoDB}
}

func (r *TransactionRepository) GetTransactionTypes(queryParams transactionResponse.TransactionTypeQuery) (response []model.TransactionType, total int64, err error) {
	coll := r.mongoDB.Client().Database(os.Getenv(constant.DATABASE_NAME)).Collection(constant.TRANSACTION_TYPES_COLLECTION)
	opts := options.Find().SetSkip(utils.GetOffsetPage(queryParams.Page, queryParams.PerPage)).SetLimit(int64(queryParams.PerPage))
	filter := bson.D{}
	if queryParams.Type != nil {
		filter = bson.D{{Key: "type", Value: queryParams.Type}}
	}

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		slog.Error("GetTransactionTypes repository: error fetching from MongoDB", "error", err)
		return nil, 0, err
	}

	var transactionTypes []model.TransactionType
	if err = cursor.All(context.TODO(), &transactionTypes); err != nil {
		slog.Error("GetTransactionTypes repository: error decode from MongoDB", "error", err)
		return nil, 0, err
	}

	count, err := coll.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		slog.Error("GetTransactionTypes repository: cannot count document", "error", err)
	}
	return transactionTypes, count, nil
}

func (r *TransactionRepository) GetTransactionsByUserId(userId string, queryParams transactionResponse.TransactionQuery) (response []transactionResponse.TransactionResponse, total int64, err error) {
	if queryParams.PerPage == 0 {
		queryParams.PerPage = 10
	}

	coll := r.mongoDB.Client().Database(os.Getenv(constant.DATABASE_NAME)).Collection(constant.TRANSACTIONS_COLLECTION)
	match := bson.D{{Key: "user_id", Value: userId}}
	if queryParams.Type != nil {
		match = append(match, bson.E{Key: "type", Value: queryParams.Type})
	}
	if queryParams.Date != nil {
		match = append(match, bson.E{Key: "date", Value: queryParams.Date})
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: match}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: constant.TRANSACTION_TYPES_COLLECTION},
			{Key: "localField", Value: "main_category_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "main_category"},
		}}},
		{{Key: "$unwind", Value: "$main_category"}},
		{{Key: "$addFields", Value: bson.D{
			{Key: "sub_category", Value: bson.D{
				{Key: "$first", Value: bson.A{
					bson.D{{
						Key: "$filter", Value: bson.D{
							{Key: "input", Value: "$main_category.sub_categories"},
							{Key: "as", Value: "sc"},
							{Key: "cond", Value: bson.D{
								{Key: "$eq", Value: bson.A{"$$sc.id", "$sub_category_id"}},
							}},
						},
					}},
				}},
			}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "user_id", Value: 1},
			{Key: "main_category_name", Value: "$main_category.name"},
			{Key: "sub_category_name", Value: "$sub_category.name"},
			{Key: "type", Value: 1},
			{Key: "amount", Value: 1},
			{Key: "date", Value: 1},
			{Key: "note", Value: 1},
		}}},
		{{Key: "$sort", Value: bson.D{{Key: "date", Value: -1}}}},
		{{Key: "$skip", Value: utils.GetOffsetPage(queryParams.Page, queryParams.PerPage)}},
		{{Key: "$limit", Value: queryParams.PerPage}},
	}

	cursor, err := coll.Aggregate(context.Background(), pipeline)
	if err != nil {
		slog.Error("GetTransactionsByUserId repository: error fetching from MongoDB", "error", err)
		return nil, 0, err
	}

	results := []transactionResponse.TransactionResponse{}
	if err = cursor.All(context.Background(), &results); err != nil {
		return results, 0, err
	}

	count, err := coll.CountDocuments(context.TODO(), match)
	if err != nil {
		slog.Error("GetTransactionsByUserId repository: cannot count document", "error", err)
	}
	return results, count, nil
}
