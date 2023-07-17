package routes

import (
	"local/witter/controllers"
	"local/witter/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProtectedRoutes(r *gin.Engine) {

	authGroup := r.Group("/auth")

	authGroup.Use(middlewares.AuthHandler("admin"))
	{
		authGroup.GET("/getmessage", controllers.GetSecretText)
	}
}
