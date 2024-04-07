package routes

import "github.com/gin-gonic/gin"

func GenerateRoutes(server *gin.Engine) {
	// Public routes :
	server.GET("/hello", hello)
	server.GET("/test", test)
	server.GET("/test2", test2)

	server.GET("/verify-token", VerifyTokenHandler)

	server.POST("/user/signup", signup)
	server.POST("/user/login", login)

	// Protected routes :
	protected := server.Group("/")
	protected.Use(AuthMiddleware())
	protected.DELETE("/user/delete", deleteUser)
	protected.POST("/chat", generalCall)
}
