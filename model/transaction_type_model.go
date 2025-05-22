package model

import "go.mongodb.org/mongo-driver/v2/bson"

type SubCategory struct {
	ID   string `bson:"id"`
	Name string
}

type TransactionType struct {
	ID            bson.ObjectID `bson:"_id"`
	Name          string
	SubCategories []SubCategory `bson:"sub_categories"`
	Type          string
}
