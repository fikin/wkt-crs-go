package cmdline

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
	assertExpectedHashFile(t, "308b4f162d916680553e219e117f1549", cfg.outFile.name)
}
