package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

// CreateDbConnection will initialize de mongo connection and gives a client
func createDbConnection() (client *mongo.Client) {

	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbhostname := os.Getenv("DB_HOSTNAME")
	uri := fmt.Sprintf("mongodb://%s:%s@%s", dbuser, dbpass, dbhostname)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	// Check the connection
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return
}

// GetClient returns
func GetClient() *mongo.Client {
	if client == nil {
		createDbConnection()
	}
	return client
}
