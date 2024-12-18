package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()          // setup an engine with logger and recovery middleware, returns pointer
	server.GET("/events", getEvents) // not executing function as a value
	server.Run(":1313")              // localhost:8080
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
