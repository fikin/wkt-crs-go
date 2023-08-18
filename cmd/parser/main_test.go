package main

import (
	// nolint:gosec
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	wktcrs "github.com/fikin/wkt-crs-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCmdOpts(t *testing.T) {
	l := log.Default()
	t.Run("missing main options", func(t *testing.T) {
		_, err := parseMainOpts(l, []string{},
			func(cwo *configWktOpt, r io.Reader, sw io.StringWriter) error {
				return fmt.Errorf("should not come here")
			},
			func(cpo *configPropsOpt) error { return fmt.Errorf("should not come here") },
		)()
		assert.ErrorContains(t, err, "pls choose \"wkt\" or \"props\" option")
	})
	t.Run("choosing -wkt, slim line", func(t *testing.T) {
		_, err := parseMainOpts(l, []string{"wkt"},
			func(cwo *configWktOpt, r io.Reader, sw io.StringWriter) error {
				assert.False(t, cwo.pretty)
				return fmt.Errorf("AAA")
			},
			func(cpo *configPropsOpt) error { return fmt.Errorf("should not come here") },
		)()
		assert.ErrorContains(t, err, "AAA")
	})
	t.Run("choosing -wkt, pretty print", func(t *testing.T) {
		_, err := parseMainOpts(l, []string{"wkt", "-pretty"},
			func(cwo *configWktOpt, r io.Reader, sw io.StringWriter) error {
				assert.True(t, cwo.pretty)
				return fmt.Errorf("AAA")
			},
			func(cpo *configPropsOpt) error { return fmt.Errorf("should not come here") },
		)()
		assert.ErrorContains(t, err, "AAA")
	})
	t.Run("choosing -props, missing package", func(t *testing.T) {
		_, err := parseMainOpts(l, []string{"props"},
			func(cwo *configWktOpt, r io.Reader, sw io.StringWriter) error {
				return fmt.Errorf("should not come here")
			},
			func(cpo *configPropsOpt) error { return fmt.Errorf("should not come here") },
		)()
		assert.ErrorContains(t, err, "define -package option")
	})
	t.Run("choosing -props, missing in", func(t *testing.T) {
		_, err := parseMainOpts(l, []string{"props", "-package", "aa"},
			func(cwo *configWktOpt, r io.Reader, sw io.StringWriter) error {
				return fmt.Errorf("should not come here")
			},
			func(cpo *configPropsOpt) error { return fmt.Errorf("should not come here") },
		)()
		assert.ErrorContains(t, err, "define -in option")
	})
	t.Run("choosing -props, missing template", func(t *testing.T) {
		_, err := parseMainOpts(l, []string{"props", "-package", "aa", "-in", "bb"},
			func(cwo *configWktOpt, r io.Reader, sw io.StringWriter) error {
				return fmt.Errorf("should not come here")
			},
			func(cpo *configPropsOpt) error { return fmt.Errorf("should not come here") },
		)()
		assert.ErrorContains(t, err, "define -templates option")
	})
	t.Run("choosing -props, missing out", func(t *testing.T) {
		_, err := parseMainOpts(l, []string{"props", "-package", "aa", "-in", "bb", "-templates", "cc"},
			func(cwo *configWktOpt, r io.Reader, sw io.StringWriter) error {
				return fmt.Errorf("should not come here")
			},
			func(cpo *configPropsOpt) error { return fmt.Errorf("should not come here") },
		)()
		assert.ErrorContains(t, err, "define -out option")
	})
	t.Run("choosing -props, ok", func(t *testing.T) {
		_, err := parseMainOpts(l, []string{"props", "-package", "aa", "-in", "bb", "-templates", "cc", "-out", "dd"},
			func(cwo *configWktOpt, r io.Reader, sw io.StringWriter) error {
				return fmt.Errorf("should not come here")
			},
			func(cpo *configPropsOpt) error {
				assert.Equal(t, "aa", cpo.packageName)
				assert.Equal(t, "bb", cpo.propsFile)
				assert.Equal(t, "cc", cpo.templateFiles)
				assert.Equal(t, "dd", cpo.outFile)
				return fmt.Errorf("AAAA")
			},
		)()
		assert.ErrorContains(t, err, "AAA")
	})
}

func TestWKT(t *testing.T) {
	cfg := &configWktOpt{
		log: log.Default(),
	}
	in := `PROJCS["NAD83(2011) / Alabama East",GEOGCS["NAD83(2011)",DATUM["NAD83 (National Spatial Reference System 2011)",SPHEROID["GRS 1980",6378137.0,298.257222101,AUTHORITY["EPSG","7019"]],AUTHORITY["EPSG","1116"]],PRIMEM["Greenwich",0.0,AUTHORITY["EPSG","8901"]],UNIT["degree",0.017453292519943295],AXIS["Geodetic longitude",EAST],AXIS["Geodetic latitude",NORTH],AUTHORITY["EPSG","6318"]],PROJECTION["Transverse_Mercator",AUTHORITY["EPSG","9807"]],PARAMETER["central_meridian",-85.83333333333333],PARAMETER["latitude_of_origin",30.499999999999996],PARAMETER["scale_factor",0.99996],PARAMETER["false_easting",200000.0],PARAMETER["false_northing",0.0],UNIT["m",1.0],AXIS["Easting",EAST],AXIS["Northing",NORTH],AUTHORITY["EPSG","6355"]]`
	t.Run("slim line", func(t *testing.T) {
		buf := &strings.Builder{}
		assert.NoError(t, handlerWkt(cfg, strings.NewReader(in), buf))
		assertExpectedHashStr(t, "28afff449ece0909917ad4188b88e080", buf.String())
	})
	t.Run("pretty print", func(t *testing.T) {
		cfg.pretty = true
		buf := &strings.Builder{}
		assert.NoError(t, handlerWkt(cfg, strings.NewReader(in), buf))
		assertExpectedHashStr(t, "6f67573f4ff4fdd60016b7c38302b330", buf.String())
	})
}

func TestTemplates(t *testing.T) {
	cfg := &configPropsOpt{
		log:           log.Default(),
		version:       "v1",
		packageName:   "wgs84",
		propsFile:     "../../testdata/v1.properties", // "../../target/geotools-epsg.properties",
		templateFiles: "../../testdata/v1*.tmpl",
		outFile:       "../../target/v1-test.txt",
	}
	assert.NoError(t, handleTemplate(cfg))
	assertExpectedHashFile(t, "4b581468f9d37cf6deb3381039caf034", cfg.outFile)
}

func TestDumpIndex(t *testing.T) {
	is, err := antlr.NewFileStream("../../testdata/v1.properties") // "../../target/geotools-epsg.properties")
	require.NoError(t, err)
	propsFileNode := wktcrs.NewWktcrsv1Parser(is).PropsFile()
	indx, err := wktcrs.AuthoritiesIndexPropFile(propsFileNode)
	require.NoError(t, err)
	data, err := json.MarshalIndent(indx, "", "  ")
	require.NoError(t, err)
	// nolint:gosec
	assert.NoError(t, os.WriteFile("../../target/index.json", data, 0o644))
	assertExpectedHashFile(t, "7ec26cb2e41362e393897ef4c0514e6d", "../../target/index.json")
}

func assertExpectedHashFile(t *testing.T, expHash, fName string) {
	// nolint:gosec
	fileContent, err := os.ReadFile(fName)
	assert.NoError(t, err)
	assertExpectedHashStr(t, expHash, string(fileContent))
}

func assertExpectedHashStr(t *testing.T, expHash, actual string) {
	// nolint:gosec
	h := md5.New()
	_, err := io.WriteString(h, actual)
	assert.NoError(t, err)
	// assert.Equal(t, exp, fmt.Sprintf("%x", h.Sum(nil)), "%s", actual)
	assert.Equal(t, expHash, fmt.Sprintf("%x", h.Sum(nil)), "%s", actual[0:int(math.Min(100, float64(len(actual))))])
}
