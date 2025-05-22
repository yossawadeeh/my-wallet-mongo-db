package response

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type TransactionTypeQuery struct {
	RequestParams
	Type *string `query:"type" validate:"oneof=INCOME OUTCOME"`
}
type SubCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type TransactionTypeResponse struct {
	ID            string                `json:"id"`
	Name          string                `json:"name"`
	SubCategories []SubCategoryResponse `json:"subCategories,omitempty"`
	Type          string                `json:"type"`
}

type TransactionQuery struct {
	RequestParams
	Type *string    `query:"type" validate:"oneof=INCOME OUTCOME"`
	Date *time.Time `query:"date" time_format:"2006-01-02"`
}
type TransactionResponse struct {
	ID               bson.ObjectID `bson:"_id" json:"id"`
	UserId           string        `bson:"user_id" json:"userId"`
	MainCategoryName string        `bson:"main_category_name" json:"mainCategoryName"`
	SubCategoryName  *string       `bson:"sub_category_name" json:"subCategoryName,omitempty"`
	Type             string        `bson:"type" json:"type"`
	Amount           float64       `bson:"amount" json:"amount"`
	Date             time.Time     `bson:"date" json:"date"`
	Note             string        `bson:"note" json:"note"`
}
