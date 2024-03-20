package routes

import (
	"milio/apicalls"
	"milio/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func generalCall(context *gin.Context) {

	var userChat models.UserChat
	err := context.ShouldBindJSON(&userChat)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}

	message := apicalls.GeneralPrompt + userChat.Message

	res, err := apicalls.CallMistralAPI(message, 5)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": "Can't get response from mistral"})
	}

	answer := res.Choices[0].Message.Content
	output := "Sorry, I can't help you with this."

	switch answer[0] {
	case '1':
		output := apicalls.SoftwareCall(userChat.Message)
		context.JSON(http.StatusOK, gin.H{"System message": output})
		return
	case '2':
		output := apicalls.HardwareCall(userChat.Message)
		context.JSON(http.StatusOK, gin.H{"System message": output})
		return
	case '3':
		output := apicalls.DomoCall(userChat.Message)
		context.JSON(http.StatusOK, gin.H{"System message": output})
		return
	case '4':
		output := apicalls.SearchCall(userChat.Message)
		context.JSON(http.StatusOK, gin.H{"System message": output})
		return
	case '5':
		output := apicalls.LogicCall(userChat.Message)
		context.JSON(http.StatusOK, gin.H{"System message": output})
		return
	case '6':
		output := apicalls.CreativeCall(userChat.Message)
		context.JSON(http.StatusOK, gin.H{"System message": output})
		return
	case '7':
		output := apicalls.CsCall(userChat.Message)
		context.JSON(http.StatusOK, gin.H{"System message": output})
		return
	}

	context.JSON(http.StatusOK, gin.H{"System message": output})
}
