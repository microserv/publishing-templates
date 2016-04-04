package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetTemplate(c *gin.Context) {
	c.JSON(404, "ads")
}

func GetAllTemplates(c *gin.Context) {
}

func AddTemplate(c *gin.Context) {
	name := c.PostForm("name")
	template := c.PostForm("template")
	fmt.Printf("Incoming template: \"%v\"\n%v\n", name, template)
}

func UpdateTemplate(c *gin.Context) {
}

func DeleteTemplate(c *gin.Context) {
}
