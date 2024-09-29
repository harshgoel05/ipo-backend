package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func initClient() *mongo.Client {
	if client != nil {
		return client
	}

	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		fmt.Errorf("failed to connect to MongoDB: %v", err)
		return nil
	}

	fmt.Println("Connected to MongoDB")

	return client
}

func getCollection(collName string) *mongo.Collection {
	client := initClient()
	return client.Database("theipolist").Collection(collName)
}
