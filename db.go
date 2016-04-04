package main

import (
    "log"
    "labix.org/v2/mgo"
    "time"
)

// Save the session here for easy access
var mongoSession *mgo.Session

// Details we need to connect to the database
// @TODO: Let's use variables from ENV, like
// os.Getenv("DB_HOST")
const (
    // Host of database to connect to
    DB_HOST = "127.0.0.1" 
    // Name of database to connect to
    DB_NAME = "templates"
    // Username of database user to connect as
    DB_USER = ""
    // Password for database user
    DB_PASS = ""
    // MongoDB Collection to use for queries
    DB_COLL = "templates"
)

type Template struct {
    Name string  // Name of the template
    Template string  // The actual template
}

// Set up the database connection.
// Stores session in the global variable "mongoSession".
func connectToDb() {
    dbConnectionInfo := &mgo.DialInfo{
        Addrs: []string{DB_HOST},
        Timeout: 10 * time.Second,
        Database: DB_NAME,
        Username: DB_USER,
        Password: DB_PASS,
    }
    
    var err error
    mongoSession, err = mgo.DialWithInfo(dbConnectionInfo)
    if err != nil {
        log.Printf("failed %v\n", err)
    }
    
    mongoSession.SetMode(mgo.Monotonic, true)
}

// Query the "templates"-collection with the given search parameters.
// Returns a []map[string]string with the results, or nil if none.
func queryCollection(searchParams map[string]string) []map[string]string {
    sessionCopy := mongoSession.Copy()
    defer sessionCopy.Close()
    
    collection := sessionCopy.DB(DB_NAME).C(DB_COLL)
    
    var results []map[string]string
    
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
    
    var _templates []map[string]string
    
    _templates = queryCollection(searchParams)

    var templates []Template

    for i, _ := range _templates {
        templates = append(templates, 
            Template{
                Name: _templates[i]["name"],
                Template: _templates[i]["template"],
            })
    }
    
    return templates
}
