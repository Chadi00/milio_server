package main

import (
	"milio/db"
	"milio/routes"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	server := gin.Default()

	db.AddErrorTable()
	db.AddUsersTable()

	routes.GenerateRoutes(server)

	server.Run()
}
