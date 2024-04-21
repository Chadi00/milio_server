package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func GenerateRoutes(server *gin.Engine) {

	v := &VisitorData{
		visitors: make(map[string]*rate.Limiter),
	}

	server.Use(RateLimitMiddleware(v)) //middleware for rate limits (2 per sec)

	server.GET("/test", test) //route to test api

	// Public routes :
	server.GET("/verify-token", VerifyTokenHandler)

	server.POST("/user/signup", signup)
	server.POST("/user/login", login)

	server.GET("/email/login", handleLogin)
	server.GET("/email/handleCallback", handleCallback)
	server.POST("/chat/stream", chatStream)

	// Protected routes :
	protected := server.Group("/")
	protected.Use(AuthMiddleware())
	protected.DELETE("/user/delete", deleteUser)
	protected.POST("/chat", generalCall)
	protected.POST("/email/get-email", getEmailAddress)
	protected.POST("/email/send", handleSendEmail)

}

func test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Message": "OK"})
}
