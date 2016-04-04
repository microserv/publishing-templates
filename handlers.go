package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "labix.org/v2/mgo"
	"strconv"
)

func generateDummyTemplate() map[string]string {
	var template = make(map[string]string)
	template["name"] = "template"
	template["template"] = "<html><body><h1>Hello world</h1><p>Some template</p></body></html>"
	return template
}

func generateDummyTemplates(num int) [10]map[string]string {
	var response [10]map[string]string
	limit := 10
	
	if num < limit {
        limit = num
    }
	
	for i := 0; i < limit; i++ {
		response[i] = generateDummyTemplate()
	}
	
	return response
}

func generateJSONErr(status_code int, message string) map[string]string {
	var response = make(map[string]string)
	response["status_code"] = strconv.Itoa(status_code)
	response["message"] = message
	return response
}

func GetTemplate(c *gin.Context) {
	c.JSON(404, "ads")
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
