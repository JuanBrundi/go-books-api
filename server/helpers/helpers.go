package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to DB
func ConnectDB() *mongo.Collection {

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

// ErrorResponse : This is error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

// Configuration model
type Configuration struct {
	Port          string
	ConnectString string
}

// GetConfiguration method basically populate configuration
// func GetConfiguration() Configuration {
// 	err := godotenv.Load("./.env")

// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	configuration := Configuration{
// 		os.Getenv("PORT"),
// 		os.Getenv("CONNECTION_STRING"),
// 	}

// 	return configuration
// }
