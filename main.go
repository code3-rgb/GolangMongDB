package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	client := Connect("mongodb://localhost:27017/myapp")
	db := client.Database("Golang")
	defer client.Disconnect(GetContext())

	collection := db.Collection("tutorial")

	result, err := collection.InsertOne(GetContext(), bson.D{
		{Key: "Third", Value: "Do something about your house"},
	})

	ErrorCheck(err)
	fmt.Println(result.InsertedID)

}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func GetContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func Connect(url string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	ErrorCheck(err)

	err = client.Connect(GetContext())

	ErrorCheck(err)

	return client

}
