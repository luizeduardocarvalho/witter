package main

import (
	"local/witter/docs"
	"local/witter/routes"
	"local/witter/shared"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// swagger embed files

var DB = make(map[string]string)

func main() {

	docs.SwaggerInfo.Title = "Witter Sample Project"

	//Db Connect and Close
	shared.Init()
	defer shared.CloseDb()

	//Init Gin
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.InitRouter(r)

	//Run Server
	r.Run(":8080")
}
