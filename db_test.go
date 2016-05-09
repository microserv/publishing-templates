package main

import (
  "fmt"
  "testing"
  
  "github.com/stretchr/testify/assert"
  )

func TestLoadTemplate(t *testing.T) {
  templates := getTemplatesByName("artikkel")  // We know this is in the db
  
  if (len(templates) > 1) {
    // wtf
  } else {
    fmt.Printf("%v", templates)
    assert.Equal(t, "artikkel", "templates[0][\"name\"]")
  }
}
