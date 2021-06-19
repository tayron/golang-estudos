package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

var result struct {
	_id  string
	Nome string
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://root:123456@172.19.0.3:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("game_of_throne").Collection("usuarios")

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		//var result2 bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
