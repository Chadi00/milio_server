package main

import (
	"milio/db"
	"milio/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
)

func main() {
	server := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))

	db.AddErrorTable()
	db.AddUsersTable()

	routes.GenerateRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}
