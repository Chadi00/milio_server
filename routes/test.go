package routes

import (
	"milio/apicalls"
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

	context.JSON(http.StatusBadRequest, gin.H{"System message": res})
}
