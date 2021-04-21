package helpers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to DB
func ConnectDD() *mongo.Collection {

	//Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

	//Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal()
	}

	fmt.Println("Connected to MongoDB")

	collection := client.Database("practice").Collection("books")

	return collection
}
