package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Transaction struct {
	ID             bson.ObjectID `bson:"_id"`
	UserId         string        `bson:"user_id"`
	MainCategoryId bson.ObjectID `bson:"main_category_id"`
	SubCategoryId  string        `bson:"sub_category_id"`
	Type           string
	Amount         float64
	Date           time.Time
	Note           string
}
