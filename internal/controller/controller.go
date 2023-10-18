package controller

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"mongoGo/internal/models"
	"mongoGo/pkg/middleware"
	"net/http"
	"strings"
)

//var client = mongo.NewMongoDatabase()

var CreateUserEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var person models.User
	err := json.NewDecoder(request.Body).Decode(&person)
	if err != nil {
		middleware.ServerErrResponse(err.Error(), response)
		return
	}
	collection := client.Database("golang").Collection("people")
	result, err := collection.InsertOne(context.TODO(), person)
	if err != nil {
		middleware.ServerErrResponse(err.Error(), response)
		return
	}
	res, err := json.Marshal(result.InsertedID)
	if err != nil {
		log.Fatal(err)
	}
	middleware.SuccessResponse(`Inserted at `+strings.Replace(string(res), `"`, ``, 2), response)
})

var GetUserEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	var person models.User

	collection := client.Database("golang").Collection("people")
	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&person)
	if err != nil {
		middleware.ErrorResponse("Person does not exist", response)
		return
	}
	middleware.SuccessRespond(person, response)
})

var GetUsersEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var people []*models.User

	collection := client.Database("golang").Collection("users")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		middleware.ServerErrResponse(err.Error(), response)
		return
	}
	for cursor.Next(context.TODO()) {
		var person models.User
		err := cursor.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		people = append(people, &person)
	}
	if err := cursor.Err(); err != nil {
		middleware.ServerErrResponse(err.Error(), response)
		return
	}
	middleware.SuccessArrRespond(people, response)
})
