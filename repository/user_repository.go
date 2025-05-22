package repository

import (
	"context"
	"errors"
	"log/slog"
	"my-wallet-ntier-mongo/constant"
	"my-wallet-ntier-mongo/interface/contract"
	"my-wallet-ntier-mongo/model"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type userRepository struct {
	mongoDB *mongo.Database
}

func NewUserRepository(mongoDB *mongo.Database) contract.UserRepository {
	return &userRepository{mongoDB: mongoDB}
}

func (r *userRepository) GetUsers() (response []model.User, total int64, err error) {
	coll := r.mongoDB.Client().Database(os.Getenv(constant.DATABASE_NAME)).Collection(constant.USERS_COLLECTION)
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		slog.Error("GetUsers repository: error fetching from MongoDB", "error", err)
		return nil, 0, err
	}

	var users []model.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		slog.Error("GetUsers repository: error decode from MongoDB", "error", err)
		return nil, 0, err
	}

	count, err := coll.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		slog.Error("GetUsers repository: cannot count document", "error", err)
	}

	return users, count, nil
}

func (r *userRepository) GetUserById(userId string) (response *model.User, err error) {
	var user model.User

	objectId, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		slog.Error("GetUser repository: Invalid ObjectID", "userId", userId, "error", err)
		return nil, errors.New(constant.INVALID_TYPE)
	}

	coll := r.mongoDB.Client().Database(os.Getenv(constant.DATABASE_NAME)).Collection(constant.USERS_COLLECTION)
	filter := bson.D{{Key: "_id", Value: objectId}}

	err = coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			slog.Error("GetUser repository: Not found", "error", err)
			return nil, errors.New(constant.DOCUMENT_NOT_FOUND)
		}
		slog.Error("GetUser repository: error fetching from MongoDB", "error", err)
		return nil, err
	}

	return &user, err
}
