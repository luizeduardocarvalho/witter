package routes

import "github.com/gin-gonic/gin"

func RegisterUtilityRoutes(r *gin.Engine) {
	registerRing(r)
}

// @Summary Ping Pong
// @Success 200 {string} string "pong"
func registerRing(r *gin.Engine) {
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
