package main

import (
	"milio/db"
	"milio/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db.AddErrorTable()
	db.AddUsersTable()

	routes.GenerateRoutes(server)

	server.Run(":8080") //localhost:8080
}
