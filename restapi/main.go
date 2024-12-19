package main

import (
	"restapi/db"
	"restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() // setup an engine with logger and recovery middleware, returns pointer
	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}
