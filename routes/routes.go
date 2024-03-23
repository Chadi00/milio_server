package routes

import "github.com/gin-gonic/gin"

func GenerateRoutes(server *gin.Engine) {
	server.POST("/chat", generalCall)

	server.GET("/hello", hello)
	server.GET("/test", test2)

	server.POST("/user/signup", signup)
	server.POST("/user/login", login)
	server.DELETE("/user/delete", deleteUser)
}
