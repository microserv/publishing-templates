package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/ory-am/dockertest.v2"
	"labix.org/v2/mgo"
	"time"

	"github.com/gin-gonic/gin"
)

func doRequest(engine *gin.Engine, method, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, nil)
	engine.ServeHTTP(w, req)
	return w
}

func TestRoutes(t *testing.T) {
	// Set up database
	var db *mgo.Session
	c, err := dockertest.ConnectToMongoDB(15, time.Millisecond*500, func(url string) bool {
		// Check if the docker image responds
		var err error
		db, err = mgo.Dial(url)
		if err != nil {
			return false
		}

		// Check if the docker image responds to ping
		return db.Ping() == nil
	})

	if err != nil {
		// log.Fatalf("Could not connect to database: %s", err)
	}

	// Close DB connection and kill container when we're done. It'll now be ready.
	defer db.Close()
	defer c.KillRemove()
	// Done setting up database

	// Set up Gin
	engine := gin.Default()

	setupServer(db)
	AddRoutes("1", engine)
	// Done setting up Gin

	var w *httptest.ResponseRecorder

	// Assert HTTP 200 is returned when loading all templates view
	w = doRequest(engine, "GET", "/api/v1/template/")
	assert.Equal(t, 200, w.Code)

	// Assert HTTP 404 if loading non-existing template
	w = doRequest(engine, "GET", "/api/v1/template/nonexisting-template")
	assert.Equal(t, 404, w.Code)

	// Load default templates
	loadDefaultTemplates("templates/")

	// Load existing article
	w = doRequest(engine, "GET", "/api/v1/template/artikkel")
	assert.Equal(t, 200, w.Code)
}
