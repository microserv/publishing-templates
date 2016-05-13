package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
)

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
func ConnectToDb() *mgo.Session {
  var mongoSession *mgo.Session
  
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
  
  return mongoSession
}

// Query the "templates"-collection with the given search parameters.
// Returns a []Template with the results, or nil if none.
func queryCollection(db *mgo.Session, searchParams map[string]string) []Template {
	sessionCopy := db.Copy()
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

func getAllTemplates(db *mgo.Session) []Template {
	sessionCopy := db.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(DB_NAME).C(DB_COLL)

	var results []Template
	collection.Find(nil).All(&results)
	return results
}

// Get all templates matching this name.
func getTemplatesByName(db *mgo.Session, name string) []Template {
	var searchParams = make(map[string]string)
	searchParams["name"] = name
  
  coll := queryCollection(db, searchParams)
  if coll != nil {
	   return coll 
  } else {
    return nil
  }
}

func insertTemplate(db *mgo.Session, template Template) error {
	sessionCopy := db.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(DB_NAME).C(DB_COLL)
	return collection.Insert(template)
}

func updateTemplate(db *mgo.Session, template Template) error {
	sessionCopy := db.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(DB_NAME).C(DB_COLL)
	return collection.Update(bson.M{"name": template.Name}, template)
}

func deleteTemplate(db *mgo.Session, template_name string) error {
	sessionCopy := db.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(DB_NAME).C(DB_COLL)
	return collection.Remove(bson.M{"name": template_name})
}
