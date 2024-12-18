package main

import (
	"net/http"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // setup an engine with logger and recovery middleware, returns pointer

	server.GET("/events", getEvents) // not executing function as a value
	server.POST("events", createEvent)

	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}
