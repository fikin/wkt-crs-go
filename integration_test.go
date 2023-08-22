//go:build integration
// +build integration

package wktcrs

import (
	"testing"

	antlr "github.com/antlr4-go/antlr/v4"
)

func TestIntegrationParsingEntireEpsgFile(t *testing.T) {
	tcData := []struct {
		name         string
		outFileName  string
		expectedHash string
		assertFn     func(t *testing.T, is antlr.CharStream) string
	}{
		{"to properties", "target/epsg.properties", "b0c6bcd0a69fd73402e84574dcff7e9c", assertToProps},
		{"to map", "target/epsg.json", "c2da8c74895cc60c7cb1d9614f042e38", assertToJSON},
		{"to pretty text", "target/epsg.txt", "3f97966a58a366230da5fbf79518d81c", assertToPrettyText},
	}
	for _, tc := range tcData {
		t.Run(tc.name, func(t *testing.T) {
			is := assertEpsgPropsFile(t)
			out := tc.assertFn(t, is)
			assertExpectedHash(t, tc.expectedHash, out)
			assertSaveToFile(t, tc.outFileName, out)
		})
	}
}
