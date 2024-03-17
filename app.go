package main

import (
	"milio/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.GenerateRoutes(server)

	server.Run(":8080") //localhost:8080
}
