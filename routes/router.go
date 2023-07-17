package routes

import (
	"local/witter/controllers"
	"local/witter/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	InitMiddleware(engine)
	authController := new(controllers.AuthController)
	engine.POST("/login", authController.HandleLogin)
	RegisterProtectedRoutes(engine)
	RegisterPublicRoutes(engine)
	RegisterUtilityRoutes(engine)
}

func InitMiddleware(engine *gin.Engine) {
	engine.Use(middlewares.CORSMiddleware())
}
