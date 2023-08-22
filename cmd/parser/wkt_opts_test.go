package main

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWKT(t *testing.T) {
	l := log.Default()
	in := `PROJCS["NAD83(2011) / Alabama East", GEOGCS["NAD83(2011)",DATUM["NAD83 (National Spatial Reference System 2011)",SPHEROID["GRS 1980",6378137.0,298.257222101,AUTHORITY["EPSG","7019"]],AUTHORITY["EPSG","1116"]],PRIMEM["Greenwich",0.0,AUTHORITY["EPSG","8901"]],UNIT["degree",0.017453292519943295],AXIS["Geodetic longitude",EAST],AXIS["Geodetic latitude",NORTH],AUTHORITY["EPSG","6318"]],PROJECTION["Transverse_Mercator",AUTHORITY["EPSG","9807"]],PARAMETER["central_meridian",-85.83333333333333],PARAMETER["latitude_of_origin",30.499999999999996],PARAMETER["scale_factor",0.99996],PARAMETER["false_easting",200000.0],PARAMETER["false_northing",0.0],UNIT["m",1.0],AXIS["Easting",EAST],AXIS["Northing",NORTH], AUTHORITY["EPSG","6355"]]`

	t.Run("usage", func(t *testing.T) {
		_, _, err := parseWktFlags(l, []string{"-dummy"})
		assert.ErrorContains(t, err, "flag provided but not defined")
	})

	for k, v := range map[string]string{
		"slim":   "45a29dceb252a9fe9babc7e80f6cc197",
		"pretty": "b5d52b3dce26dc18a2688c633e1fe239",
		"json":   "92268834f354d15c74a67b6e108d5b31",
	} {
		t.Run(k, func(t *testing.T) {
			cfg, _, err := parseWktFlags(l, []string{"-format", k})
			require.NoError(t, err)
			buf := &bytes.Buffer{}
			cfg.inFile.file = strings.NewReader(in)
			cfg.outFile.file = buf
			require.NoError(t, handlerWkt(cfg))
			assertExpectedHashStr(t, v, buf.String())
		})
	}
}
