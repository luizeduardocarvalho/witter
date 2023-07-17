package routes

import (
	"local/witter/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUtilityRoutes(r *gin.Engine) {
	registerRing(r)
}

func registerRing(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)
}
