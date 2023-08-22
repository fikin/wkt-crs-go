package cmdline

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPropsOpts(t *testing.T) {
	l := log.Default()

	t.Run("usage", func(t *testing.T) {
		_, _, err := parseWktFlags(l, []string{"-dummy"})
		assert.ErrorContains(t, err, "flag provided but not defined")
	})

	for k, v := range map[string]string{
		"index":  "916d4895e3ab492daefdfbf506eb04ad",
		"json":   "1928581679758d16b55f0fcfd4140e10",
		"pretty": "e0e394ceb6434428abfe697064d03a18",
		"slim":   "95ba3d9928e094ecaa1ffecf20a16a33",
		"wgs84":  "673892a25d92a36a1f38e74f95fa6ce6",
	} {
		t.Run(k, func(t *testing.T) {
			cfg, _, err := parsePropsFlags(l, []string{"-in", "../../testdata/v1.properties", "-out=-", "-format", k})
			require.NoError(t, err)
			buf := bytes.Buffer{}
			cfg.outFile.file = &buf
			require.NoError(t, handleProps(cfg))
			assertExpectedHashStr(t, v, buf.String())
		})
	}
}
