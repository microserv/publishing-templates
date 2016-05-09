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

	loadDefaultTemplates("templates/")
	engine.Run(":8002")
}

const enable_access_control = false

func AddRoutes(version string, engine *gin.Engine) {
	oauth2_routes := engine.Group("oauth2")
	{
		oauth2_routes.GET("/callback", OAuth2Callback)
	}

	public_routes := engine.Group(fmt.Sprintf("api/v%s", version))
	{
		public_routes.GET("/template/", GetAllTemplates)
		public_routes.GET("/template/:template_name", GetTemplate)
	}

	template_restricted := engine.Group(fmt.Sprintf("api/v%s", version))
	if enable_access_control {
        template_restricted.Use(ValidateRequest("write"))
    }
	{
		template_restricted.POST("/template/", InsertTemplate)
		template_restricted.PUT("/template/:template_name", GetTemplate)
		template_restricted.DELETE("/template/:template_name", DeleteTemplate)
	}
}
