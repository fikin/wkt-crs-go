package wktcrs

import (
	"bytes"
	"math"

	// nolint:gosec
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	antlr "github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const r1 = `
TEST2=
	GEOGCS["WGS 84", 
		DATUM["World Geodetic System 1984", 
			SPHEROID["WGS 84", 6378137.0, 298.257223563, 
				AUTHORITY["EPSG","7030"]
			]
		], 
		PRIMEM["Greenwich", 0.0, 
			AUTHORITY["EPSG","8901"]
		], 
		UNIT["degree", 0.017453292519943295], 
		AXIS["Longitude", EAST], 
		AXIS["Latitude", NORTH], 
		AUTHORITY["EPSG","TEST2"]
	]
`

func TestParsingAConstant(t *testing.T) {
	tcData := []struct {
		name         string
		expectedHash string
		assertFn     func(t *testing.T, is antlr.CharStream) string
	}{
		{"to properties", "3347d8bc7d5af93731edf01fd6d81a9e", assertToProps},
		{"to map", "fe3322c331932b8d19da4f452ae14167", assertToJSON},
		{"to pretty text", "e9cd386c52c7f442d3de9833cc987d95", assertToPrettyText},
	}
	for _, tc := range tcData {
		t.Run(tc.name, func(t *testing.T) {
			is := antlr.NewInputStream(r1)
			out := tc.assertFn(t, is)
			assertExpectedHash(t, tc.expectedHash, out)
		})
	}
}

func TestParsingEpsgFile(t *testing.T) {
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

func assertExpectedHash(t *testing.T, exp, actual string) {
	// nolint:gosec
	h := md5.New()
	_, err := io.WriteString(h, actual)
	assert.NoError(t, err)
	// assert.Equal(t, exp, fmt.Sprintf("%x", h.Sum(nil)), "%s", actual)
	assert.Equal(t, exp, fmt.Sprintf("%x", h.Sum(nil)), "%s", actual[0:int(math.Min(100, float64(len(actual))))])
}

func assertToProps(t *testing.T, is antlr.CharStream) string {
	tree := ParsePropsFileAST(is)
	out := &strings.Builder{}
	assert.NoError(t, NodeToText(tree, out))
	return out.String()
}

func assertToJSON(t *testing.T, is antlr.CharStream) string {
	tree := ParsePropsFileAST(is)
	arr, err := PropsFileNodeToArr(tree)
	require.NoError(t, err)
	out := &bytes.Buffer{}
	ee := json.NewEncoder(out)
	ee.SetIndent("", "  ")
	assert.NoError(t, ee.Encode(arr))
	return out.String()
}

func assertToPrettyText(t *testing.T, is antlr.CharStream) string {
	tree := ParsePropsFileAST(is)
	out := &strings.Builder{}
	assert.NoError(t, NodeToPrettyText(tree, "", "  ", out))
	return out.String()
}

func assertEpsgPropsFile(t *testing.T) (is antlr.CharStream) {
	is, err := antlr.NewFileStream("target/geotools-epsg.properties")
	require.NoError(t, err)
	return is
}

func assertSaveToFile(t *testing.T, oFileName, data string) {
	// nolint:gosec
	fo, err := os.Create(oFileName)
	require.NoError(t, err)
	_, err = fo.WriteString(data)
	assert.NoError(t, err)
	require.NoError(t, fo.Close())
}
