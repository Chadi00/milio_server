package main

import (
	"milio/db"
	"milio/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db.AddErrorTable()
	db.AddUsersTable()

	routes.GenerateRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}
