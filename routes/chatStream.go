package routes

import (
	"log"
	"milio/apicalls"
	"milio/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func chatStream(context *gin.Context) {
	var chat models.UserChat
	if err := context.BindJSON(&chat); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing message content"})
		return
	}

	if chat.Message == "" {
		log.Println("No message content received")
		context.JSON(http.StatusBadRequest, gin.H{"error": "No message content provided"})
		return
	}

	if err := apicalls.CallAnthropicAPIStreaming(context, chat.Message, 2000, 0.5); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process message"})
	}
}
