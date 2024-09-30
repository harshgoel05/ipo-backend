package main

import (
	"context"
	"log"
	"time"
)

func fetchIpoDetailsAndInsertInDb(ipoList []DMIPO) []SMIPOIndividual {
	var ipoDetailsFinal []SMIPOIndividual
	for _, ipo := range ipoList {
		ipoDetails := fetchIndividualIpoDetails(ipo.Link, ipo.GmpUrl)
		ipoDetailsResponse := SMIPOIndividual{
			Slug:        ipo.Slug,
			Details:     ipoDetails.Details,
			GmpTimeline: ipoDetails.GmpTimeline,
		}
		// Insert in database
		insertIpoDetailsInDatabase(ipoDetailsResponse)
		ipoDetailsFinal = append(ipoDetailsFinal, ipoDetailsResponse)
		time.Sleep(2 * time.Second)
	}
	return ipoDetailsFinal
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

func updateGmpAndDetailsForAllIposIndividual(slug string) []AMIPOIndividual {

	aMIPOIndividual := fetchIpoDetailsFromDatabase(slug)
	println("aMIPOIndividual", aMIPOIndividual)
	dmIPO := convertAMIPOIndividualToDMIPO(*aMIPOIndividual)
	ipoDetailsFinal := fetchIpoDetailsAndInsertInDb([]DMIPO{dmIPO})
	response := mapIpoBasicInfoToDetailedInfoBySlug([]DMIPO{dmIPO}, ipoDetailsFinal)
	return response
}

func mapIpoBasicInfoToDetailedInfoBySlug(ipoList []DMIPO, ipoDetails []SMIPOIndividual) []AMIPOIndividual {
	var ipoDetailsFinal []AMIPOIndividual
	for _, ipo := range ipoList {
		for _, ipoDetail := range ipoDetails {
			if ipo.Slug == ipoDetail.Slug {

				temp := mergeDMIPOAndSMIPOIndividualToAMIPOIndividual(ipo, ipoDetail)
				ipoDetailsFinal = append(ipoDetailsFinal, temp)
				break
			}
		}
	}
	return ipoDetailsFinal
}
