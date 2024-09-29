package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the IPO service!",
		})
	})

	router.GET("/update/calendar", func(ctx *gin.Context) {
		ipoDetails := fetchAndUpdateCalendar()
		ctx.IndentedJSON(200, ipoDetails)
	})

	router.GET("/update/gmp-and-details", func(ctx *gin.Context) {
		ipoDetails := updateGmpAndDetailsForAllIpos()
		ctx.IndentedJSON(200, ipoDetails)
	})

	router.GET("/calendar", func(ctx *gin.Context) {
		ipoDetails := fetchCalendarFromDatabase()
		ctx.IndentedJSON(200, ipoDetails)
	})

	// router.GET("/details/:slug", getIpoDetails)
	router.Run()
}
