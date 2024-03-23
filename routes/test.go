package routes

import (
	"milio/apicalls"
	"milio/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(context *gin.Context) {
	message := "Hello! tell me a joke please."

	res, err := apicalls.CallMistralAPI(message, 5)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": "Can't get response from mistral"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"System message": res})
}

func test2(context *gin.Context) {
	res, err := db.ReadAllErrorsAsString()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": "error with readAllErrorsAsString"})
	}

	context.JSON(http.StatusOK, gin.H{"errors": res})
}
