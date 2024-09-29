package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func getIpoCalendarFromCrawl() []DMIPO {
	url := BASE_URL + CALENDAR_API

	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal the response body into the Event struct
	var ipoList []DMIPO
	err = json.Unmarshal(body, &ipoList)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return ipoList
}

func fetchAndUpdateCalendar() []DMIPO {
	docs := getIpoCalendarFromCrawl()
	operations := updateOrInsertIPOCalendar(docs)
	collection := getCollection("ipo_calendar")
	// Execute bulk write
	result, err := collection.BulkWrite(context.TODO(), operations)
	if err != nil {
		log.Fatalf("Bulk write failed: %v", err)
	}
	log.Printf("Bulk write result: %v", result)
	return docs

}
