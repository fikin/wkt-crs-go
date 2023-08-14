package wktcrs

import (
	"bytes"
	// nolint:gosec
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"testing"

	antlr "github.com/antlr4-go/antlr/v4"
	"github.com/fikin/wkt-crs-go/wktcrsv1"
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
		{"to map", "target/epsg.json", "8dc2a7c5689b767c50b06c208c7c6c93", assertToJSON},
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
	assert.Equal(t, exp, fmt.Sprintf("%x", h.Sum(nil)), "%s", actual[0:int(math.Min(100, float64(len(actual))))])
}

func parseTree(is antlr.CharStream) wktcrsv1.IPropsFileContext {
	lexer := wktcrsv1.Newwktcrsv1Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := wktcrsv1.Newwktcrsv1Parser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	return parser.PropsFile()
}

func assertToProps(t *testing.T, is antlr.CharStream) string {
	tree := parseTree(is)
	out := &strings.Builder{}
	assert.NoError(t, nodeToText(tree, out))
	return out.String()
}

func assertToJSON(t *testing.T, is antlr.CharStream) string {
	tree := parseTree(is)
	arr, err := propsFileNodeToArr(tree)
	require.NoError(t, err)
	out := &bytes.Buffer{}
	ee := json.NewEncoder(out)
	ee.SetIndent("", "  ")
	assert.NoError(t, ee.Encode(arr))
	return out.String()
}

func assertToPrettyText(t *testing.T, is antlr.CharStream) string {
	tree := parseTree(is)
	out := &strings.Builder{}
	assert.NoError(t, nodeToPrettyText(tree, "", "  ", out))
	return out.String()
}

func assertEpsgPropsFile(t *testing.T) (is antlr.CharStream) {
	is, err := antlr.NewFileStream("target/epsg.properties")
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
