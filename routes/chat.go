package routes

import (
	"log"
	"milio/apicalls"
	"milio/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func analyzeChat(context *gin.Context) {
	var userChat models.UserChat

	err := context.ShouldBindJSON(&userChat)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("MIXTRAL_API_KEY")
	if apiKey == "" {
		log.Fatal("API key not set in .env file")
	}

	message := userChat.Message
	sysChat, err := apicalls.CallMistralAPI(message, apiKey)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get a response back from Mixtral"})
	}

	context.JSON(http.StatusOK, gin.H{"System message": sysChat.Choices[0].Message.Content})

}
