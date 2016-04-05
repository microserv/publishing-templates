package main

import (
	"labix.org/v2/mgo"
	"log"
	"time"
)

// Save the session here for easy access
var mongoSession *mgo.Session

// Details we need to connect to the database
// @TODO: Let's use variables from ENV, like
// os.Getenv("DB_HOST")
const (
	DB_HOST = "127.0.0.1"
	DB_NAME = "templates"
	DB_USER = ""
	DB_PASS = ""
	DB_COLL = "templates"
)

type Template struct {
	Name string "json:'name'"
	Html string "json:'html'"
}

// Set up the database connection.
// Stores session in the global variable "mongoSession".
func connectToDb() {
	dbConnectionInfo := &mgo.DialInfo{
		Addrs:    []string{DB_HOST},
		Timeout:  10 * time.Second,
		Database: DB_NAME,
		Username: DB_USER,
		Password: DB_PASS,
	}

	var err error
	mongoSession, err = mgo.DialWithInfo(dbConnectionInfo)
	if err != nil {
		log.Printf("Error when attempting to connect to the DB! ERROR: %v\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)
}

// Query the "templates"-collection with the given search parameters.
// Returns a []Template with the results, or nil if none.
func queryCollection(searchParams map[string]string) []Template {
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(DB_NAME).C(DB_COLL)

	var results []Template
	err := collection.Find(searchParams).All(&results)

	if err != nil {
		log.Printf("RunQuery : ERROR : %s\n", err)
		return nil
	}

	return results
}

// Get all templates matching this name.
func getTemplatesByName(name string) []Template {
	var searchParams = make(map[string]string)
	searchParams["name"] = name
	return queryCollection(searchParams)
}

func addTemplate(template Template) error {
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(DB_NAME).C(DB_COLL)
	return collection.Insert(template)
}
