package main

import (
	"net/http"

	"jwt-auth/pkg/config"
	"jwt-auth/pkg/controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	config.SetUpDB()
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to OTP verification")
	})

	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/signup", controller.SignUp)
		userRoutes.POST("/signin", controller.SignIn)
		userRoutes.POST("/verify", controller.Verify)
	}

	router.Run(":8080")
}
