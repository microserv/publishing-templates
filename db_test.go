package main

import (
	"gopkg.in/ory-am/dockertest.v2"
	"labix.org/v2/mgo"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadTemplate(t *testing.T) {
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

	// We're ready, let's test

	// We have no templates in the DB, so a query should return nothing
	no_templates := getAllTemplates(db)
	assert.Empty(t, no_templates)

	// There are no templates in the DB, so a query should return nothing
	query_templates := getTemplatesByName(db, "nothing")
	assert.Nil(t, query_templates)

	// Test inserting a template
	var template_to_insert Template
	template_to_insert.Name = "test-insert"
	template_to_insert.Html = "<h1>Hello world</h1>"
	insertTemplate(db, template_to_insert)
	templates_after_insert := getTemplatesByName(db, "test-insert")
	assert.Equal(t, 1, len(templates_after_insert))
	assert.Equal(t, template_to_insert, templates_after_insert[0])

	// Test getting all templates when we have some data
	all_templates_is_one := getAllTemplates(db)
	assert.Equal(t, 1, len(all_templates_is_one))

	// Test updating a template
	var template_to_update Template
	template_to_update.Name = "test-update"
	template_to_update.Html = "<h1>Hello world</h1>"
	insertTemplate(db, template_to_update)
	templates_after_insert_2 := getTemplatesByName(db, "test-update")
	assert.Equal(t, 1, len(templates_after_insert_2))
	assert.Equal(t, template_to_update, templates_after_insert_2[0])

	// Test getting all templates when we have some data
	all_templates_is_two := getAllTemplates(db)
	assert.Equal(t, 2, len(all_templates_is_two))

	// Notice the "!" at the end
	template_to_update.Html = "<h1>Hello world!</h1>"
	updateTemplate(db, template_to_update)
	templates_after_update := getTemplatesByName(db, "test-update")
	assert.Equal(t, 1, len(templates_after_update))
	assert.Equal(t, template_to_update, templates_after_update[0])

	// Test getting all templates when we have some more data
	all_templates_is_still_two := getAllTemplates(db)
	assert.Equal(t, 2, len(all_templates_is_still_two))

	// Test deleting a non-existing template
	template_to_delete := "test-article"
	deleteTemplate(db, template_to_delete)
	templates_after_delete_nonexist := getTemplatesByName(db, "test-article")
	assert.Equal(t, 0, len(templates_after_delete_nonexist))
	assert.Empty(t, templates_after_delete_nonexist)

	// Test deleting an existing template
	template_to_delete = "test-insert"
	deleteTemplate(db, template_to_delete)
	templates_after_delete_exist := getTemplatesByName(db, "test-insert")
	assert.Equal(t, 0, len(templates_after_delete_exist))
	assert.Empty(t, templates_after_delete_exist)

	// Test getting all templates when we have deleted
	all_templates_is_one_again := getAllTemplates(db)
	assert.Equal(t, 1, len(all_templates_is_one_again))
}
