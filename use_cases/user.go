package use_cases

import (
	"context"
	"go-api/dtos"
	"go-api/services"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindByName(name string) []dtos.User {
	collection := services.GetUserCollection()
	filter := bson.D{{"name", name}}
	cur, err := collection.Find(context.Background(), filter, options.Find())
	if err != nil {
		log.Fatal(err)
	}
	var results []dtos.User

	for cur.Next(context.Background()) {

		// create a value into which the single document can be decoded
		var elem dtos.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}
	return results
}
