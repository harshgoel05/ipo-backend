package main

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Document struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func updateOrInsertIPOCalendar(docs interface{}) []mongo.WriteModel {
	var operations []mongo.WriteModel
	val := reflect.ValueOf(docs)

	// Check if docs is a slice
	if val.Kind() != reflect.Slice {
		return operations // Return empty if not a slice
	}

	// Iterate through each element in the slice
	for i := 0; i < val.Len(); i++ {
		doc := val.Index(i)

		// Create a BSON update document
		update := bson.D{
			{Key: "$set", Value: bson.D{}},
		}

		// Use reflection to access fields
		for j := 0; j < doc.NumField(); j++ {
			field := doc.Type().Field(j)
			fieldValue := doc.Field(j)

			// Skip unexported fields
			if !fieldValue.CanInterface() {
				continue
			}

			// Get the JSON tag
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				continue // Skip if no JSON tag
			}

			// Add field to update document using JSON tag
			update[0].Value = append(update[0].Value.(bson.D), bson.E{Key: jsonTag, Value: fieldValue.Interface()})
		}

		// Create the update operation
		operation := mongo.NewUpdateOneModel().
			SetFilter(bson.D{{Key: "slug", Value: doc.FieldByName("Slug").Interface()}}). // Using "slug" as the unique identifier
			SetUpdate(update).
			SetUpsert(true)
		operations = append(operations, operation)
	}

	return operations
}

func ReadAllDocuments(collection *mongo.Collection) ([]DMIPO, error) {
	var results []DMIPO

	// Find all documents in the collection
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Iterate through the cursor and decode each document into results
	for cursor.Next(context.TODO()) {
		var doc DMIPO
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		results = append(results, doc)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// FetchIPOWithDetails performs an optimized aggregation using $lookup
func FetchIPOWithDetails(collection *mongo.Collection) ([]AMIPOIndividual, error) {
	// MongoDB aggregation pipeline
	pipeline := mongo.Pipeline{
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "ipo_details"},
				{Key: "localField", Value: "slug"},
				{Key: "foreignField", Value: "slug"},
				{Key: "pipeline", Value: bson.A{
					bson.D{
						{Key: "$project", Value: bson.D{
							{Key: "_id", Value: 0},
							{Key: "details", Value: 1},
							{Key: "gmptimeline", Value: 1},
						}},
					},
				}},
				{Key: "as", Value: "ipo_details"},
			}},
		},
		// $addFields to handle no match case
		{
			{Key: "$addFields", Value: bson.D{
				{Key: "ipo_details", Value: bson.D{
					{Key: "$cond", Value: bson.A{
						bson.D{{Key: "$eq", Value: bson.A{"$ipo_details", bson.A{}}}},
						nil,
						bson.D{{Key: "$arrayElemAt", Value: bson.A{"$ipo_details", 0}}}, // In case there is a match, take the first element
					}},
				}},
			}},
		},
		// $addFields stage for details and gmpTimeline
		{
			{Key: "$addFields", Value: bson.D{
				{Key: "details", Value: "$ipo_details.details"},
				{Key: "gmpTimeline", Value: "$ipo_details.gmptimeline"},
			}},
		},
		// $project stage to remove ipo_details
		{
			{Key: "$project", Value: bson.D{
				{Key: "ipo_details", Value: 0},
			}},
		},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []AMIPOIndividual
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func FetchIndividualIPOWithDetails(collection *mongo.Collection, slug string) (AMIPOIndividual, error) {
	// MongoDB aggregation pipeline
	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.D{{Key: "slug", Value: slug}}},
		},
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "ipo_details"},
				{Key: "localField", Value: "slug"},
				{Key: "foreignField", Value: "slug"},
				{Key: "pipeline", Value: bson.A{
					bson.D{
						{Key: "$project", Value: bson.D{
							{Key: "_id", Value: 0},
							{Key: "details", Value: 1},
							{Key: "gmptimeline", Value: 1},
						}},
					},
				}},
				{Key: "as", Value: "ipo_details"},
			}},
		},
		// $addFields to handle no match case
		{
			{Key: "$addFields", Value: bson.D{
				{Key: "ipo_details", Value: bson.D{
					{Key: "$cond", Value: bson.A{
						bson.D{{Key: "$eq", Value: bson.A{"$ipo_details", bson.A{}}}},
						nil,
						bson.D{{Key: "$arrayElemAt", Value: bson.A{"$ipo_details", 0}}}, // In case there is a match, take the first element
					}},
				}},
			}},
		},
		// $addFields stage for details and gmpTimeline
		{
			{Key: "$addFields", Value: bson.D{
				{Key: "details", Value: "$ipo_details.details"},
				{Key: "gmpTimeline", Value: "$ipo_details.gmptimeline"},
			}},
		},
		// $project stage to remove ipo_details
		{
			{Key: "$project", Value: bson.D{
				{Key: "ipo_details", Value: 0},
			}},
		},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return AMIPOIndividual{}, err
	}
	defer cursor.Close(context.TODO())

	var result AMIPOIndividual
	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&result); err != nil {
			return AMIPOIndividual{}, err
		}
	} else {
		return AMIPOIndividual{}, mongo.ErrNoDocuments
	}

	return result, nil
}
