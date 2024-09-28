package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func readIpoJson() []DMIPO {
	var ipoList []DMIPO
	jsonFile, err := os.Open("ipo_data.json")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	if err := json.Unmarshal(byteValue, &ipoList); err != nil {
		log.Fatalf("Error unmarshaling JSON: %s", err)
	}
	return ipoList

}

// getAlbums responds with the list of all albums as JSON.
func getIpoCalendar(c *gin.Context) {

	ipoDetails := readIpoJson()
	c.IndentedJSON(http.StatusOK, ipoDetails)
}

func findIpoBySlug(ipoDetails []DMIPO, slug string) (*DMIPO, bool) {
	for _, ipoDetail := range ipoDetails {
		if ipoDetail.Slug == slug {
			return &ipoDetail, true
		}
	}
	return nil, false
}

func getIpoDetails(c *gin.Context) {
	slug := c.Param("slug")
	ipoList, found := findIpoBySlug(readIpoJson(), slug)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "IPO not found"})
	} else {
		c.IndentedJSON(http.StatusOK, ipoList)
	}
}

func main() {

	router := gin.Default()
	router.GET("/calendar", getIpoCalendar)
	router.GET("/details/:slug", getIpoDetails)
	router.Run("localhost:8080")
}
