package controllers

import "github.com/gin-gonic/gin"

func GetSecretText(c *gin.Context) {
	c.JSON(200, "Hi this is a secret message. Auth was successful!")
}

// @Summary      Shows a string
// @Description  Returns a sample string
// @Tags         SampleController
// @Accept       json
// @Produce      json
// @Success      200  {object} string
// @Router       /publicmessage [get]
func GetPublicText(c *gin.Context) {
	c.JSON(200, "Hi this is a public message!")
}

// @Summary      Returns a PONG
// @Description  Returns a PONG
// @Tags         SampleController
// @Accept       json
// @Produce      json
// @Success      200  {object} string
// @Router       /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, "Pong")
}
