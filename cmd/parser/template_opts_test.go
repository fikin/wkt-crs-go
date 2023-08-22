package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTemplates(t *testing.T) {
	cfg, _, err := parseTemplateFlags(log.Default(), []string{
		"-package=wgs84",
		"-template", "../../templates/wgs84*.tmpl",
		"-in", "../../testdata/v1.properties",
		"-out", "../../target/wgs84.txt",
		"-format=wgs84",
	})
	require.NoError(t, err)
	require.NoError(t, handleTemplate(cfg))
	assertExpectedHashFile(t, "b5e56c291dfec47d03411ae130ee2f29", cfg.outFile.name)
}
