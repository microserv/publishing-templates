package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	version := "1"

	engine := gin.Default()
	AddRoutes(version, engine)

	connectToDb()

	engine.Run(":8080")
}

func AddRoutes(version string, engine *gin.Engine) {
	api_group := engine.Group(fmt.Sprintf("api/v%s", version))
	{
		api_group.GET("/template/", GetAllTemplates)
		api_group.GET("/template/:template_name", GetTemplate)

		api_group.POST("/template", AddTemplate)
		api_group.PUT("/template/", UpdateTemplate)

		api_group.DELETE("/template/:template_name", DeleteTemplate)
	}
}
