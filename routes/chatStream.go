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

	log.Printf("Received message: '%s'", chat.Message)

	res, err := apicalls.CallGroqAPI(chat.Message, 8000, 0.7)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process message"})
	}
	context.JSON(http.StatusOK, gin.H{"Response": res.Choices[0].Message.Content})
}
