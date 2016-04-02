package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Template struct {
	id   int64  "db:'id' json:'id'"
	name string "db:'name' json:'template_name'"
	html string "db:'html' json:'html'"
}

func main() {
	version := "1"
	var router *gin.Engine
	router = CreateRouter(version)
	router.Run(":8080")
}

func CreateRouter(version string) *gin.Engine {
	router := gin.Default()

	api_group := router.Group(fmt.Sprintf("api/v%s", version))
	{
		api_group.GET("/template/", GetAllTemplates)
		api_group.GET("/template/:template_name", GetTemplate)

		api_group.POST("/template", AddTemplate)
		api_group.PUT("/template/:template_name", UpdateTemplate)

		api_group.DELETE("/template/:template_name", DeleteTemplate)
	}

	return router
}

func GetTemplate(c *gin.Context) {
	c.JSON(404, "ads")
}

func GetAllTemplates(c *gin.Context) {
}

func AddTemplate(c *gin.Context) {
}

func UpdateTemplate(c *gin.Context) {
}

func DeleteTemplate(c *gin.Context) {
}
