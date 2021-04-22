package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"server/helpers"
	"server/models"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = helpers.ConnectDB()

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Book array
	var books []models.Book

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helpers.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var book models.Book

		err := cur.Decode(&book)

		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(books)
}
