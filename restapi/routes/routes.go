package routes

import (
	"restapi/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // not executing function as a value
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
