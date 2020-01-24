package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

// FibonacciDto -
type FibonacciDto struct {
	N     float64 `json:"n"`
	Value float64 `json:"value"`
}

// User entity Dto
type User struct {
	ID                   primitive.ObjectID `bson:"_id"`
	Name                 string             `bson:"name"`
	LastName             string             `bson:"lastName"`
	Gender               string             `bson:"gender"`
	Country              string             `bson:"country"`
	ProgrammingLanguages []string           `bson:"programming_languages"`
}
