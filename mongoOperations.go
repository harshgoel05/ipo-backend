package main

import (
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
