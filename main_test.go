package main

import "testing"

func TestGenerateDummyTemplate(t *testing.T) {
	dummy_template := generateDummyTemplate()

	if dummy_template["name"] != "template" || dummy_template["template"] != "<html><body><h1>Hello world</h1><p>Some template</p></body></html>" {
		t.Error("Dummy template didnt return what it should.")
	}
}

func TestGenerateJsonErr(t *testing.T) {
	json_response := generateJSONErr(200, "Testing")

	if json_response["status_code"] != "200" || json_response["message"] != "Testing" {
		t.Error("Generating Json Error failed")
	}
}
