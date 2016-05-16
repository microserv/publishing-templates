package main

import (
	"fmt"
    "time"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
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

	engine.Use(cors.Middleware(cors.Config{
		Origins:         "http://localhost:8000, https://gallifrey.sklirg.io, http://despina.128.no, https://despina.128.no",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	AddRoutes(version, engine)
	engine.Static("/static", "./static")

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
