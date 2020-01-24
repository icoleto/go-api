package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

// FibonacciDto -
type FibonacciDto struct {
	N     float64 `json:"n"`
	Value float64 `json:"value"`
}

// User entity Dto
type User struct {
	ID                   primitive.ObjectID `json:"id" bson:"_id"`
	Name                 string             `json:"name" bson:"name"`
	LastName             string             `json:"lastName" bson:"lastname"`
	Gender               string             `json:"gender" bson:"gender"`
	Country              string             `json:"country" bson:"country"`
	ProgrammingLanguages []string           `json:"programmingLanguages" bson:"programming_languages"`
}
