package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	FirstName  string             `bson:"first_name"`
	LastName   string             `bson:"last_name"`
	Privileges []string
	Avatar     string
	Email      string
	Password   string
	BirthDate  string `bson:"birth_date"`
}
