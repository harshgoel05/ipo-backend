package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func getIpoCalendarFromCrawl() []DMIPO {
	url := BASE_URL + CALENDAR_API
	// Read the response body
	body := apiCall(url)

	// Unmarshal the response body into the Event struct
	var ipoList []DMIPO
	err := json.Unmarshal(body, &ipoList)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return ipoList
}

func fetchIndividualIpoDetails(detailsUrl string, gmpUrl *string) DMIPOIndividual {
	url := BASE_URL + "/details?details_url=" + detailsUrl
	if gmpUrl != nil {
		url = url + "&gmp_url=" + *gmpUrl
	}

	// Read the response body
	body := apiCall(url)

	// Unmarshal the response body into the Event struct
	var ipo DMIPOIndividual
	err := json.Unmarshal(body, &ipo)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return ipo
}

func apiCall(url string) []byte {
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

	return body
}
