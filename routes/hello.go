package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "hello"})
}
