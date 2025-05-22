package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Address struct {
	Line1       string `bson:"line_1"`
	SubDistrict string `bson:"sub_district"`
	District    string
	Province    string
	Postcode    string
}

type User struct {
	ID        bson.ObjectID `bson:"_id"`
	FirstName string        `bson:"first_name"`
	LastName  string        `bson:"last_name"`
	Email     string
	Phone     string
	Address   Address
}
