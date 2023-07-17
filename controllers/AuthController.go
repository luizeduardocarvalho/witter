package controllers

import (
	"local/witter/middlewares"
	"local/witter/shared"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (this AuthController) HandleLogin(c *gin.Context) {
	userId := "123"
	username := "Beast"
	roles := []string{shared.RoleAdmin, shared.RoleProUser}

	// do user auth here

	//issue token
	token, err := middlewares.GenerateToken([]byte(middlewares.SigningKey), userId, username, roles)
	if err != nil {

	}
	c.JSON(200, token)
}
