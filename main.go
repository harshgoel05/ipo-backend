package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the IPO service!",
		})
	})

	router.GET("/calendar", func(ctx *gin.Context) {
		ipoDetails := fetchAndUpdateCalendar()
		ctx.IndentedJSON(200, ipoDetails)
	})
	// router.GET("/details/:slug", getIpoDetails)
	router.Run("localhost:8080")
}
