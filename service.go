package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
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

func fetchIpoDetailsAndInsertInDb(ipoList []DMIPO) []SMIPOIndividual {
	var ipoDetailsFinal []SMIPOIndividual
	for _, ipo := range ipoList {
		ipoDetails := fetchIndividualIpoDetails(ipo.Link, ipo.GmpUrl)
		ipoDetailsResponse := SMIPOIndividual{
			Slug:        ipo.Slug,
			Details:     ipoDetails.Details,
			GmpTimeline: ipoDetails.GmpTimeline,
		}
		ipoDetailsFinal = append(ipoDetailsFinal, ipoDetailsResponse)
		// Insert in database
		collection := getCollection("ipo_details")
		// Execute bulk write
		result, err := collection.InsertOne(
			context.TODO(),
			ipoDetailsResponse,
		)
		if err != nil {
			log.Fatalf("Bulk write failed: %v", err)
		}
		log.Printf("Bulk write result: %v", result)
		time.Sleep(2 * time.Second)
	}
	return ipoDetailsFinal
}

func fetchIndividualIpoDetails(detailsUrl string, gmpUrl string) DMIPOIndividual {
	url := BASE_URL + "/details?details_url=" + detailsUrl
	if gmpUrl != "" {
		url = url + "&gmp_url=" + gmpUrl
	}

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
	var ipo DMIPOIndividual
	err = json.Unmarshal(body, &ipo)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return ipo
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

func updateGmpAndDetailsForAllIpos() []AMIPOIndividual {

	collection := getCollection("ipo_calendar")
	ipoList, error := ReadAllDocuments(collection)
	ipoDetailsFinal := fetchIpoDetailsAndInsertInDb(ipoList)
	response := mapIpoBasicInfoToDetailedInfoBySlug(ipoList, ipoDetailsFinal)
	if error != nil {
		log.Fatalf("Failed to read documents: %v", error)
	}
	return response
}

func mapIpoBasicInfoToDetailedInfoBySlug(ipoList []DMIPO, ipoDetails []SMIPOIndividual) []AMIPOIndividual {
	var ipoDetailsFinal []AMIPOIndividual
	for _, ipo := range ipoList {
		for _, ipoDetail := range ipoDetails {
			if ipo.Slug == ipoDetail.Slug {
				temp := AMIPOIndividual{
					StartDate:   ipo.StartDate,
					GmpUrl:      ipo.GmpUrl,
					Link:        ipo.Link,
					EndDate:     ipo.EndDate,
					LogoUrl:     ipo.LogoUrl,
					ListingDate: ipo.ListingDate,
					PriceRange:  ipo.PriceRange,
					Symbol:      ipo.Symbol,
					Name:        ipo.Name,
					Slug:        ipo.Slug,
					Details:     ipoDetail.Details,
					GmpTimeline: ipoDetail.GmpTimeline,
				}
				ipoDetailsFinal = append(ipoDetailsFinal, temp)
				break
			}
		}
	}
	return ipoDetailsFinal
}

func fetchCalendarFromDatabase() []AMIPOIndividual {

	collection := getCollection("ipo_calendar")
	res, err := FetchIPOWithDetails(collection)
	if err != nil {
		log.Fatalf("Failed to fetch documents: %v", err)
	}
	return res
}
