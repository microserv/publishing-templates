package main

import (
	"strconv"
)

// Generates a boring dummy template.
func generateDummyTemplate() map[string]string {
	var template = make(map[string]string)
	template["name"] = "template"
	template["template"] = "<html><body><h1>Hello world</h1><p>Some template</p></body></html>"
	return template
}

// Generates num (int) boring dummy templates.
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

// Tool for generating a JSON blob containing an error and a HTTP status code
func generateJSONErr(status_code int, message string) map[string]string {
	var response = make(map[string]string)
	response["status_code"] = strconv.Itoa(status_code)
	response["message"] = message
	return response
}
