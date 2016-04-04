package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// The format for loglines
const LOGFORMAT = "[%v] %v - %v | %v | %v | %v"

func GetTemplate(c *gin.Context) {
    templates := getTemplatesByName(c.Param("template_name"))
    if len(templates) > 0 {
        c.JSON(200, templates)
    } else {
        c.JSON(404, generateJSONErr(404, "Not found"))
    }
}

func GetAllTemplates(c *gin.Context) {
	var _num = c.DefaultQuery("limit", "10")
	
	num, err := strconv.Atoi(_num)
	if err != nil {
		str := fmt.Sprintf("Invalid value for \"limit\", should be int but was %T (%v).", _num, _num)
		c.JSON(400, generateJSONErr(400, str))
		c.Abort()
		return
	}
	
	var response [10]map[string]string
	
	response = generateDummyTemplates(num)
	
	c.JSON(200, response)
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
