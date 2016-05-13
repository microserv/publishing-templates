package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetTemplate(c *gin.Context) {
	templates := getTemplatesByName(DB, c.Param("template_name"))
	if len(templates) > 0 {
		c.JSON(200, templates)
	} else {
		c.JSON(404, generateJSONErr(404, "Not found"))
	}
}

func GetAllTemplates(c *gin.Context) {
	c.JSON(200, getAllTemplates(DB))
}

func InsertTemplate(c *gin.Context) {
	// Use the JSON fields in the 'Template' struct to automatically bind the
	// JSON POST request fields to an instance of the 'Template' struct.
	var template Template
	c.Bind(&template)

	err := insertTemplate(DB, template)
	if err != nil {
		error_msg := fmt.Sprintf("An error occured while attempting to insert: %+v. ERROR: %v", template, err)
		c.JSON(400, generateJSONErr(400, error_msg))
		c.Abort()
	} else {
		c.JSON(201, "The template was successfully inserted.")
	}
}

func UpdateTemplate(c *gin.Context) {
	// Use the JSON fields in the 'Template' struct to automatically bind the
	// JSON POST request fields to an instance of the 'Template' struct.
	var template Template
	c.Bind(&template)

	err := updateTemplate(DB, template)
	if err != nil {
		error_msg := fmt.Sprintf("An error occured while attempting to update: %+v. ERROR: %v", template, err)
		c.JSON(400, generateJSONErr(400, error_msg))
		c.Abort()
	} else {
		c.JSON(200, "The template was successfully updated.")
	}
}

func DeleteTemplate(c *gin.Context) {
	template_name := c.Param("template_name")

	err := deleteTemplate(DB, template_name)
	if err != nil {
		error_msg := fmt.Sprintf("An error occured while attempting to delete: %v. ERROR: %v", template_name, err)
		c.JSON(400, generateJSONErr(400, error_msg))
		c.Abort()
	} else {
		c.JSON(204, "The template was successfully deleted.")
	}
}

func OAuth2Callback(c *gin.Context) {
	response := make(map[string]string)
	code := c.DefaultQuery("code", "")
	next := c.DefaultQuery("next", "")

	response["code"] = code
	response["next"] = next

	if code != "" {
		valid, err := ValidateUserToken(code)
		if err != "" {
			fmt.Printf("Authentication failed: %v\n", err)
		} else {
			if valid {
				// Debug
				// fmt.Println("Authentication succeeded")
			} else {
				fmt.Println("Token does not carry correct scope for this operation.")
			}
		}
		c.JSON(200, response)
		// Redirect to "next" param, as that's where the user wants to go
	} else {
		c.JSON(400, generateJSONErr(400, "Missing code"))
	}
}
