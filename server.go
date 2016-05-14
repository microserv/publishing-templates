package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"labix.org/v2/mgo"
)

var DB *mgo.Session

func GetDbConnection() *mgo.Session {
	return DB
}

func main() {
	version := "1"

	setupServer(nil)

	if DB == nil {
		fmt.Printf("DB connection failed: %s", DB)
	}

	engine := gin.Default()
	AddRoutes(version, engine)

	loadDefaultTemplates("templates/")
	engine.Run()
}

func setupServer(db *mgo.Session) {
	if db == nil {
		DB = ConnectToDb()
	} else {
		DB = db
	}
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
