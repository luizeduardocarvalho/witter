package routes

import (
	"local/witter/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.Engine) {

	r.GET("/publicmessage", controllers.GetPublicText)
}
