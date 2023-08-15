package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplates(t *testing.T) {
	cfg := config{
		log:           log.Default(),
		version:       "v1",
		packageName:   "test",
		propsFile:     "../../testdata/v1.properties",
		templateFiles: "../../testdata/v1*.tmpl",
		outFile:       "../../target/v1-test.txt",
	}
	assert.NoError(t, handleTemplate(cfg))
}
