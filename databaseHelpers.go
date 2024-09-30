package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func fetchIpoDetailsFromDatabase(slug string) *AMIPOIndividual {

	collection := getCollection("ipo_calendar")
	res, err := FetchIndividualIPOWithDetails(collection, slug)
	if err != nil {
		print("Error in fetching ipo details for slug: ", slug)
		return nil
	}
	return &res
}

func fetchCalendarFromDatabase() []AMIPOIndividual {

	collection := getCollection("ipo_calendar")
	res, err := FetchIPOWithDetails(collection)
	if err != nil {
		print("Error in fetching calendar from database", err)
		return []AMIPOIndividual{}
	}
	return res
}

func insertIpoDetailsInDatabase(sMIPOIndividual SMIPOIndividual) {
	collection := getCollection("ipo_details")

	filter := bson.M{"slug": sMIPOIndividual.Slug}

	update := bson.M{
		"$set": DMIPOIndividual{
			Details:     sMIPOIndividual.Details,
			GmpTimeline: sMIPOIndividual.GmpTimeline,
		},
	}

	// Perform upsert
	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(
		context.TODO(),
		filter,
		update,
		opts,
	)
	if err != nil {
		log.Fatalf("Upsert failed: %v", err)
	}
	log.Printf("Upsert result: %v - %v ", sMIPOIndividual.Slug, result)
}
