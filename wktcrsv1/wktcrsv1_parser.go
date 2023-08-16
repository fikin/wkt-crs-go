// Code generated from wktcrsv1.g4 by ANTLR 4.13.0. DO NOT EDIT.

package wktcrsv1 // wktcrsv1
import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type wktcrsv1Parser struct {
	*antlr.BaseParser
}

var Wktcrsv1ParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func wktcrsv1ParserInit() {
  staticData := &Wktcrsv1ParserStaticData
  staticData.LiteralNames = []string{
    "", "'COMPD_CS'", "'PROJCS'", "'GEOCCS'", "'GEOGCS'", "'VERT_CS'", "'LOCAL_CS'", 
    "'DATUM'", "'VERT_DATUM'", "'LOCAL_DATUM'", "'SPHEROID'", "'TOWGS84'", 
    "'AUTHORITY'", "'PRIMEM'", "'UNIT'", "'AXIS'", "'PROJECTION'", "'PARAMETER'", 
    "'\"EPSG\"'", "'\"ESRI\"'", "'EAST'", "'WEST'", "'NORTH'", "'SOUTH'", 
    "'NORTH_EAST'", "'NORTH_WEST'", "'UP'", "'DOWN'", "'GEOCENTRIC_X'", 
    "'GEOCENTRIC_Y'", "'GEOCENTRIC_Z'", "", "", "", "", "", "','", "", "", 
    "'='",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "NUMBER", "TEXT", 
    "PKEY", "COMMENT_LINE", "WS", "COMMA", "LPAR", "RPAR", "EQ",
  }
  staticData.RuleNames = []string{
    "propsFile", "propRow", "commentLine", "epsgDefLine", "wkt", "compdcs", 
    "projcs", "geoccs", "geogcs", "vertcs", "localcs", "datum", "vertdatum", 
    "localdatum", "spheroid", "towgs84", "authority", "primem", "unit", 
    "axis", "projection", "parameter", "authorityName", "axisOrient", "epsgCode", 
    "name", "number", "type", "semiMajorAxis", "inverseFlattening", "dx", 
    "dy", "dz", "ex", "ey", "ez", "ppm", "code", "longitude", "conversionFactor", 
    "value",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 39, 395, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2, 
	21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26, 
	7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7, 
	31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36, 
	2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 5, 0, 84, 
	8, 0, 10, 0, 12, 0, 87, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 93, 8, 1, 1, 
	2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 
	4, 107, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 115, 8, 5, 1, 5, 
	1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 
	1, 6, 1, 6, 1, 6, 1, 6, 4, 6, 134, 8, 6, 11, 6, 12, 6, 135, 1, 6, 1, 6, 
	1, 6, 1, 6, 1, 6, 5, 6, 143, 8, 6, 10, 6, 12, 6, 146, 9, 6, 1, 6, 1, 6, 
	1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 
	1, 7, 1, 7, 4, 7, 164, 8, 7, 11, 7, 12, 7, 165, 1, 7, 1, 7, 1, 7, 1, 8, 
	1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 182, 
	8, 8, 10, 8, 12, 8, 185, 9, 8, 1, 8, 1, 8, 3, 8, 189, 8, 8, 1, 8, 1, 8, 
	1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 
	1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 
	10, 1, 10, 4, 10, 217, 8, 10, 11, 10, 12, 10, 218, 1, 10, 1, 10, 1, 10, 
	1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 231, 8, 11, 1, 
	11, 1, 11, 3, 11, 235, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 
	1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 
	13, 1, 13, 3, 13, 255, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 
	1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 268, 8, 14, 1, 14, 1, 14, 1, 
	15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 
	1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 287, 8, 15, 3, 15, 289, 8, 15, 1, 15, 
	1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 
	17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 307, 8, 17, 1, 17, 1, 17, 1, 18, 
	1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 318, 8, 18, 1, 18, 1, 
	18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 
	1, 20, 1, 20, 3, 20, 334, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 
	21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 
	1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 359, 8, 23, 1, 
	24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 
	1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 
	34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 
	1, 40, 1, 40, 1, 40, 0, 0, 41, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 
	60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 0, 2, 1, 0, 18, 19, 2, 0, 31, 
	31, 33, 33, 387, 0, 85, 1, 0, 0, 0, 2, 92, 1, 0, 0, 0, 4, 94, 1, 0, 0, 
	0, 6, 96, 1, 0, 0, 0, 8, 106, 1, 0, 0, 0, 10, 108, 1, 0, 0, 0, 12, 122, 
	1, 0, 0, 0, 14, 150, 1, 0, 0, 0, 16, 170, 1, 0, 0, 0, 18, 192, 1, 0, 0, 
	0, 20, 205, 1, 0, 0, 0, 22, 223, 1, 0, 0, 0, 24, 238, 1, 0, 0, 0, 26, 247, 
	1, 0, 0, 0, 28, 258, 1, 0, 0, 0, 30, 271, 1, 0, 0, 0, 32, 292, 1, 0, 0, 
	0, 34, 299, 1, 0, 0, 0, 36, 310, 1, 0, 0, 0, 38, 321, 1, 0, 0, 0, 40, 328, 
	1, 0, 0, 0, 42, 337, 1, 0, 0, 0, 44, 344, 1, 0, 0, 0, 46, 358, 1, 0, 0, 
	0, 48, 360, 1, 0, 0, 0, 50, 362, 1, 0, 0, 0, 52, 364, 1, 0, 0, 0, 54, 366, 
	1, 0, 0, 0, 56, 368, 1, 0, 0, 0, 58, 370, 1, 0, 0, 0, 60, 372, 1, 0, 0, 
	0, 62, 374, 1, 0, 0, 0, 64, 376, 1, 0, 0, 0, 66, 378, 1, 0, 0, 0, 68, 380, 
	1, 0, 0, 0, 70, 382, 1, 0, 0, 0, 72, 384, 1, 0, 0, 0, 74, 386, 1, 0, 0, 
	0, 76, 388, 1, 0, 0, 0, 78, 390, 1, 0, 0, 0, 80, 392, 1, 0, 0, 0, 82, 84, 
	3, 2, 1, 0, 83, 82, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 
	85, 86, 1, 0, 0, 0, 86, 88, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 89, 5, 
	0, 0, 1, 89, 1, 1, 0, 0, 0, 90, 93, 3, 4, 2, 0, 91, 93, 3, 6, 3, 0, 92, 
	90, 1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 3, 1, 0, 0, 0, 94, 95, 5, 34, 0, 
	0, 95, 5, 1, 0, 0, 0, 96, 97, 3, 48, 24, 0, 97, 98, 5, 39, 0, 0, 98, 99, 
	3, 8, 4, 0, 99, 7, 1, 0, 0, 0, 100, 107, 3, 10, 5, 0, 101, 107, 3, 12, 
	6, 0, 102, 107, 3, 16, 8, 0, 103, 107, 3, 18, 9, 0, 104, 107, 3, 14, 7, 
	0, 105, 107, 3, 20, 10, 0, 106, 100, 1, 0, 0, 0, 106, 101, 1, 0, 0, 0, 
	106, 102, 1, 0, 0, 0, 106, 103, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 
	105, 1, 0, 0, 0, 107, 9, 1, 0, 0, 0, 108, 109, 5, 1, 0, 0, 109, 110, 5, 
	37, 0, 0, 110, 111, 3, 50, 25, 0, 111, 114, 5, 36, 0, 0, 112, 115, 3, 12, 
	6, 0, 113, 115, 3, 16, 8, 0, 114, 112, 1, 0, 0, 0, 114, 113, 1, 0, 0, 0, 
	115, 116, 1, 0, 0, 0, 116, 117, 5, 36, 0, 0, 117, 118, 3, 18, 9, 0, 118, 
	119, 5, 36, 0, 0, 119, 120, 3, 32, 16, 0, 120, 121, 5, 38, 0, 0, 121, 11, 
	1, 0, 0, 0, 122, 123, 5, 2, 0, 0, 123, 124, 5, 37, 0, 0, 124, 125, 3, 50, 
	25, 0, 125, 126, 5, 36, 0, 0, 126, 127, 3, 16, 8, 0, 127, 128, 5, 36, 0, 
	0, 128, 129, 3, 40, 20, 0, 129, 133, 5, 36, 0, 0, 130, 131, 3, 42, 21, 
	0, 131, 132, 5, 36, 0, 0, 132, 134, 1, 0, 0, 0, 133, 130, 1, 0, 0, 0, 134, 
	135, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 137, 
	1, 0, 0, 0, 137, 138, 3, 36, 18, 0, 138, 144, 5, 36, 0, 0, 139, 140, 3, 
	38, 19, 0, 140, 141, 5, 36, 0, 0, 141, 143, 1, 0, 0, 0, 142, 139, 1, 0, 
	0, 0, 143, 146, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 
	145, 147, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 147, 148, 3, 32, 16, 0, 148, 
	149, 5, 38, 0, 0, 149, 13, 1, 0, 0, 0, 150, 151, 5, 3, 0, 0, 151, 152, 
	5, 37, 0, 0, 152, 153, 3, 50, 25, 0, 153, 154, 5, 36, 0, 0, 154, 155, 3, 
	22, 11, 0, 155, 156, 5, 36, 0, 0, 156, 157, 3, 34, 17, 0, 157, 158, 5, 
	36, 0, 0, 158, 159, 3, 36, 18, 0, 159, 163, 5, 36, 0, 0, 160, 161, 3, 38, 
	19, 0, 161, 162, 5, 36, 0, 0, 162, 164, 1, 0, 0, 0, 163, 160, 1, 0, 0, 
	0, 164, 165, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 
	167, 1, 0, 0, 0, 167, 168, 3, 32, 16, 0, 168, 169, 5, 38, 0, 0, 169, 15, 
	1, 0, 0, 0, 170, 171, 5, 4, 0, 0, 171, 172, 5, 37, 0, 0, 172, 173, 3, 50, 
	25, 0, 173, 174, 5, 36, 0, 0, 174, 175, 3, 22, 11, 0, 175, 176, 5, 36, 
	0, 0, 176, 177, 3, 34, 17, 0, 177, 178, 5, 36, 0, 0, 178, 183, 3, 36, 18, 
	0, 179, 180, 5, 36, 0, 0, 180, 182, 3, 38, 19, 0, 181, 179, 1, 0, 0, 0, 
	182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 
	188, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 187, 5, 36, 0, 0, 187, 189, 
	3, 32, 16, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 190, 1, 
	0, 0, 0, 190, 191, 5, 38, 0, 0, 191, 17, 1, 0, 0, 0, 192, 193, 5, 5, 0, 
	0, 193, 194, 5, 37, 0, 0, 194, 195, 3, 50, 25, 0, 195, 196, 5, 36, 0, 0, 
	196, 197, 3, 24, 12, 0, 197, 198, 5, 36, 0, 0, 198, 199, 3, 36, 18, 0, 
	199, 200, 5, 36, 0, 0, 200, 201, 3, 38, 19, 0, 201, 202, 5, 36, 0, 0, 202, 
	203, 3, 32, 16, 0, 203, 204, 5, 38, 0, 0, 204, 19, 1, 0, 0, 0, 205, 206, 
	5, 6, 0, 0, 206, 207, 5, 37, 0, 0, 207, 208, 3, 50, 25, 0, 208, 209, 5, 
	36, 0, 0, 209, 210, 3, 26, 13, 0, 210, 211, 5, 36, 0, 0, 211, 212, 3, 36, 
	18, 0, 212, 216, 5, 36, 0, 0, 213, 214, 3, 38, 19, 0, 214, 215, 5, 36, 
	0, 0, 215, 217, 1, 0, 0, 0, 216, 213, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 
	218, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 
	221, 3, 32, 16, 0, 221, 222, 5, 38, 0, 0, 222, 21, 1, 0, 0, 0, 223, 224, 
	5, 7, 0, 0, 224, 225, 5, 37, 0, 0, 225, 226, 3, 50, 25, 0, 226, 227, 5, 
	36, 0, 0, 227, 230, 3, 28, 14, 0, 228, 229, 5, 36, 0, 0, 229, 231, 3, 30, 
	15, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 234, 1, 0, 0, 0, 
	232, 233, 5, 36, 0, 0, 233, 235, 3, 32, 16, 0, 234, 232, 1, 0, 0, 0, 234, 
	235, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 237, 5, 38, 0, 0, 237, 23, 
	1, 0, 0, 0, 238, 239, 5, 8, 0, 0, 239, 240, 5, 37, 0, 0, 240, 241, 3, 50, 
	25, 0, 241, 242, 5, 36, 0, 0, 242, 243, 3, 54, 27, 0, 243, 244, 5, 36, 
	0, 0, 244, 245, 3, 32, 16, 0, 245, 246, 5, 38, 0, 0, 246, 25, 1, 0, 0, 
	0, 247, 248, 5, 9, 0, 0, 248, 249, 5, 37, 0, 0, 249, 250, 3, 50, 25, 0, 
	250, 251, 5, 36, 0, 0, 251, 254, 3, 54, 27, 0, 252, 253, 5, 36, 0, 0, 253, 
	255, 3, 32, 16, 0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 
	1, 0, 0, 0, 256, 257, 5, 38, 0, 0, 257, 27, 1, 0, 0, 0, 258, 259, 5, 10, 
	0, 0, 259, 260, 5, 37, 0, 0, 260, 261, 3, 50, 25, 0, 261, 262, 5, 36, 0, 
	0, 262, 263, 3, 56, 28, 0, 263, 264, 5, 36, 0, 0, 264, 267, 3, 58, 29, 
	0, 265, 266, 5, 36, 0, 0, 266, 268, 3, 32, 16, 0, 267, 265, 1, 0, 0, 0, 
	267, 268, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 270, 5, 38, 0, 0, 270, 
	29, 1, 0, 0, 0, 271, 272, 5, 11, 0, 0, 272, 273, 5, 37, 0, 0, 273, 274, 
	3, 60, 30, 0, 274, 275, 5, 36, 0, 0, 275, 276, 3, 62, 31, 0, 276, 277, 
	5, 36, 0, 0, 277, 288, 3, 64, 32, 0, 278, 279, 5, 36, 0, 0, 279, 280, 3, 
	66, 33, 0, 280, 281, 5, 36, 0, 0, 281, 282, 3, 68, 34, 0, 282, 283, 5, 
	36, 0, 0, 283, 286, 3, 70, 35, 0, 284, 285, 5, 36, 0, 0, 285, 287, 3, 72, 
	36, 0, 286, 284, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 289, 1, 0, 0, 0, 
	288, 278, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 
	291, 5, 38, 0, 0, 291, 31, 1, 0, 0, 0, 292, 293, 5, 12, 0, 0, 293, 294, 
	5, 37, 0, 0, 294, 295, 3, 44, 22, 0, 295, 296, 5, 36, 0, 0, 296, 297, 3, 
	74, 37, 0, 297, 298, 5, 38, 0, 0, 298, 33, 1, 0, 0, 0, 299, 300, 5, 13, 
	0, 0, 300, 301, 5, 37, 0, 0, 301, 302, 3, 50, 25, 0, 302, 303, 5, 36, 0, 
	0, 303, 306, 3, 76, 38, 0, 304, 305, 5, 36, 0, 0, 305, 307, 3, 32, 16, 
	0, 306, 304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 
	309, 5, 38, 0, 0, 309, 35, 1, 0, 0, 0, 310, 311, 5, 14, 0, 0, 311, 312, 
	5, 37, 0, 0, 312, 313, 3, 50, 25, 0, 313, 314, 5, 36, 0, 0, 314, 317, 3, 
	78, 39, 0, 315, 316, 5, 36, 0, 0, 316, 318, 3, 32, 16, 0, 317, 315, 1, 
	0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 320, 5, 38, 0, 
	0, 320, 37, 1, 0, 0, 0, 321, 322, 5, 15, 0, 0, 322, 323, 5, 37, 0, 0, 323, 
	324, 3, 50, 25, 0, 324, 325, 5, 36, 0, 0, 325, 326, 3, 46, 23, 0, 326, 
	327, 5, 38, 0, 0, 327, 39, 1, 0, 0, 0, 328, 329, 5, 16, 0, 0, 329, 330, 
	5, 37, 0, 0, 330, 333, 3, 50, 25, 0, 331, 332, 5, 36, 0, 0, 332, 334, 3, 
	32, 16, 0, 333, 331, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 335, 1, 0, 
	0, 0, 335, 336, 5, 38, 0, 0, 336, 41, 1, 0, 0, 0, 337, 338, 5, 17, 0, 0, 
	338, 339, 5, 37, 0, 0, 339, 340, 3, 50, 25, 0, 340, 341, 5, 36, 0, 0, 341, 
	342, 3, 80, 40, 0, 342, 343, 5, 38, 0, 0, 343, 43, 1, 0, 0, 0, 344, 345, 
	7, 0, 0, 0, 345, 45, 1, 0, 0, 0, 346, 359, 5, 20, 0, 0, 347, 359, 5, 21, 
	0, 0, 348, 359, 5, 22, 0, 0, 349, 359, 5, 23, 0, 0, 350, 359, 5, 24, 0, 
	0, 351, 359, 5, 25, 0, 0, 352, 359, 5, 26, 0, 0, 353, 359, 5, 27, 0, 0, 
	354, 359, 5, 28, 0, 0, 355, 359, 5, 29, 0, 0, 356, 359, 5, 30, 0, 0, 357, 
	359, 3, 50, 25, 0, 358, 346, 1, 0, 0, 0, 358, 347, 1, 0, 0, 0, 358, 348, 
	1, 0, 0, 0, 358, 349, 1, 0, 0, 0, 358, 350, 1, 0, 0, 0, 358, 351, 1, 0, 
	0, 0, 358, 352, 1, 0, 0, 0, 358, 353, 1, 0, 0, 0, 358, 354, 1, 0, 0, 0, 
	358, 355, 1, 0, 0, 0, 358, 356, 1, 0, 0, 0, 358, 357, 1, 0, 0, 0, 359, 
	47, 1, 0, 0, 0, 360, 361, 7, 1, 0, 0, 361, 49, 1, 0, 0, 0, 362, 363, 5, 
	32, 0, 0, 363, 51, 1, 0, 0, 0, 364, 365, 5, 31, 0, 0, 365, 53, 1, 0, 0, 
	0, 366, 367, 5, 31, 0, 0, 367, 55, 1, 0, 0, 0, 368, 369, 5, 31, 0, 0, 369, 
	57, 1, 0, 0, 0, 370, 371, 5, 31, 0, 0, 371, 59, 1, 0, 0, 0, 372, 373, 5, 
	31, 0, 0, 373, 61, 1, 0, 0, 0, 374, 375, 5, 31, 0, 0, 375, 63, 1, 0, 0, 
	0, 376, 377, 5, 31, 0, 0, 377, 65, 1, 0, 0, 0, 378, 379, 5, 31, 0, 0, 379, 
	67, 1, 0, 0, 0, 380, 381, 5, 31, 0, 0, 381, 69, 1, 0, 0, 0, 382, 383, 5, 
	31, 0, 0, 383, 71, 1, 0, 0, 0, 384, 385, 5, 31, 0, 0, 385, 73, 1, 0, 0, 
	0, 386, 387, 5, 32, 0, 0, 387, 75, 1, 0, 0, 0, 388, 389, 5, 31, 0, 0, 389, 
	77, 1, 0, 0, 0, 390, 391, 5, 31, 0, 0, 391, 79, 1, 0, 0, 0, 392, 393, 5, 
	31, 0, 0, 393, 81, 1, 0, 0, 0, 20, 85, 92, 106, 114, 135, 144, 165, 183, 
	188, 218, 230, 234, 254, 267, 286, 288, 306, 317, 333, 358,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// wktcrsv1ParserInit initializes any static state used to implement wktcrsv1Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// Newwktcrsv1Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Wktcrsv1ParserInit() {
  staticData := &Wktcrsv1ParserStaticData
  staticData.once.Do(wktcrsv1ParserInit)
}

// Newwktcrsv1Parser produces a new parser instance for the optional input antlr.TokenStream.
func Newwktcrsv1Parser(input antlr.TokenStream) *wktcrsv1Parser {
	Wktcrsv1ParserInit()
	this := new(wktcrsv1Parser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &Wktcrsv1ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "wktcrsv1.g4"

	return this
}


// wktcrsv1Parser tokens.
const (
	wktcrsv1ParserEOF = antlr.TokenEOF
	wktcrsv1ParserT__0 = 1
	wktcrsv1ParserT__1 = 2
	wktcrsv1ParserT__2 = 3
	wktcrsv1ParserT__3 = 4
	wktcrsv1ParserT__4 = 5
	wktcrsv1ParserT__5 = 6
	wktcrsv1ParserT__6 = 7
	wktcrsv1ParserT__7 = 8
	wktcrsv1ParserT__8 = 9
	wktcrsv1ParserT__9 = 10
	wktcrsv1ParserT__10 = 11
	wktcrsv1ParserT__11 = 12
	wktcrsv1ParserT__12 = 13
	wktcrsv1ParserT__13 = 14
	wktcrsv1ParserT__14 = 15
	wktcrsv1ParserT__15 = 16
	wktcrsv1ParserT__16 = 17
	wktcrsv1ParserT__17 = 18
	wktcrsv1ParserT__18 = 19
	wktcrsv1ParserT__19 = 20
	wktcrsv1ParserT__20 = 21
	wktcrsv1ParserT__21 = 22
	wktcrsv1ParserT__22 = 23
	wktcrsv1ParserT__23 = 24
	wktcrsv1ParserT__24 = 25
	wktcrsv1ParserT__25 = 26
	wktcrsv1ParserT__26 = 27
	wktcrsv1ParserT__27 = 28
	wktcrsv1ParserT__28 = 29
	wktcrsv1ParserT__29 = 30
	wktcrsv1ParserNUMBER = 31
	wktcrsv1ParserTEXT = 32
	wktcrsv1ParserPKEY = 33
	wktcrsv1ParserCOMMENT_LINE = 34
	wktcrsv1ParserWS = 35
	wktcrsv1ParserCOMMA = 36
	wktcrsv1ParserLPAR = 37
	wktcrsv1ParserRPAR = 38
	wktcrsv1ParserEQ = 39
)

// wktcrsv1Parser rules.
const (
	wktcrsv1ParserRULE_propsFile = 0
	wktcrsv1ParserRULE_propRow = 1
	wktcrsv1ParserRULE_commentLine = 2
	wktcrsv1ParserRULE_epsgDefLine = 3
	wktcrsv1ParserRULE_wkt = 4
	wktcrsv1ParserRULE_compdcs = 5
	wktcrsv1ParserRULE_projcs = 6
	wktcrsv1ParserRULE_geoccs = 7
	wktcrsv1ParserRULE_geogcs = 8
	wktcrsv1ParserRULE_vertcs = 9
	wktcrsv1ParserRULE_localcs = 10
	wktcrsv1ParserRULE_datum = 11
	wktcrsv1ParserRULE_vertdatum = 12
	wktcrsv1ParserRULE_localdatum = 13
	wktcrsv1ParserRULE_spheroid = 14
	wktcrsv1ParserRULE_towgs84 = 15
	wktcrsv1ParserRULE_authority = 16
	wktcrsv1ParserRULE_primem = 17
	wktcrsv1ParserRULE_unit = 18
	wktcrsv1ParserRULE_axis = 19
	wktcrsv1ParserRULE_projection = 20
	wktcrsv1ParserRULE_parameter = 21
	wktcrsv1ParserRULE_authorityName = 22
	wktcrsv1ParserRULE_axisOrient = 23
	wktcrsv1ParserRULE_epsgCode = 24
	wktcrsv1ParserRULE_name = 25
	wktcrsv1ParserRULE_number = 26
	wktcrsv1ParserRULE_type = 27
	wktcrsv1ParserRULE_semiMajorAxis = 28
	wktcrsv1ParserRULE_inverseFlattening = 29
	wktcrsv1ParserRULE_dx = 30
	wktcrsv1ParserRULE_dy = 31
	wktcrsv1ParserRULE_dz = 32
	wktcrsv1ParserRULE_ex = 33
	wktcrsv1ParserRULE_ey = 34
	wktcrsv1ParserRULE_ez = 35
	wktcrsv1ParserRULE_ppm = 36
	wktcrsv1ParserRULE_code = 37
	wktcrsv1ParserRULE_longitude = 38
	wktcrsv1ParserRULE_conversionFactor = 39
	wktcrsv1ParserRULE_value = 40
)

// IPropsFileContext is an interface to support dynamic dispatch.
type IPropsFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllPropRow() []IPropRowContext
	PropRow(i int) IPropRowContext

	// IsPropsFileContext differentiates from other interfaces.
	IsPropsFileContext()
}

type PropsFileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropsFileContext() *PropsFileContext {
	var p = new(PropsFileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_propsFile
	return p
}

func InitEmptyPropsFileContext(p *PropsFileContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_propsFile
}

func (*PropsFileContext) IsPropsFileContext() {}

func NewPropsFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropsFileContext {
	var p = new(PropsFileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_propsFile

	return p
}

func (s *PropsFileContext) GetParser() antlr.Parser { return s.parser }

func (s *PropsFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserEOF, 0)
}

func (s *PropsFileContext) AllPropRow() []IPropRowContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropRowContext); ok {
			len++
		}
	}

	tst := make([]IPropRowContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropRowContext); ok {
			tst[i] = t.(IPropRowContext)
			i++
		}
	}

	return tst
}

func (s *PropsFileContext) PropRow(i int) IPropRowContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropRowContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropRowContext)
}

func (s *PropsFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropsFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) PropsFile() (localctx IPropsFileContext) {
	localctx = NewPropsFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, wktcrsv1ParserRULE_propsFile)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 27917287424) != 0) {
		{
			p.SetState(82)
			p.PropRow()
		}


		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(88)
		p.Match(wktcrsv1ParserEOF)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IPropRowContext is an interface to support dynamic dispatch.
type IPropRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CommentLine() ICommentLineContext
	EpsgDefLine() IEpsgDefLineContext

	// IsPropRowContext differentiates from other interfaces.
	IsPropRowContext()
}

type PropRowContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropRowContext() *PropRowContext {
	var p = new(PropRowContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_propRow
	return p
}

func InitEmptyPropRowContext(p *PropRowContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_propRow
}

func (*PropRowContext) IsPropRowContext() {}

func NewPropRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropRowContext {
	var p = new(PropRowContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_propRow

	return p
}

func (s *PropRowContext) GetParser() antlr.Parser { return s.parser }

func (s *PropRowContext) CommentLine() ICommentLineContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentLineContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentLineContext)
}

func (s *PropRowContext) EpsgDefLine() IEpsgDefLineContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsgDefLineContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsgDefLineContext)
}

func (s *PropRowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropRowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) PropRow() (localctx IPropRowContext) {
	localctx = NewPropRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, wktcrsv1ParserRULE_propRow)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case wktcrsv1ParserCOMMENT_LINE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.CommentLine()
		}


	case wktcrsv1ParserNUMBER, wktcrsv1ParserPKEY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.EpsgDefLine()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ICommentLineContext is an interface to support dynamic dispatch.
type ICommentLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMENT_LINE() antlr.TerminalNode

	// IsCommentLineContext differentiates from other interfaces.
	IsCommentLineContext()
}

type CommentLineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentLineContext() *CommentLineContext {
	var p = new(CommentLineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_commentLine
	return p
}

func InitEmptyCommentLineContext(p *CommentLineContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_commentLine
}

func (*CommentLineContext) IsCommentLineContext() {}

func NewCommentLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentLineContext {
	var p = new(CommentLineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_commentLine

	return p
}

func (s *CommentLineContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentLineContext) COMMENT_LINE() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMENT_LINE, 0)
}

func (s *CommentLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) CommentLine() (localctx ICommentLineContext) {
	localctx = NewCommentLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, wktcrsv1ParserRULE_commentLine)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(wktcrsv1ParserCOMMENT_LINE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IEpsgDefLineContext is an interface to support dynamic dispatch.
type IEpsgDefLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EpsgCode() IEpsgCodeContext
	EQ() antlr.TerminalNode
	Wkt() IWktContext

	// IsEpsgDefLineContext differentiates from other interfaces.
	IsEpsgDefLineContext()
}

type EpsgDefLineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEpsgDefLineContext() *EpsgDefLineContext {
	var p = new(EpsgDefLineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_epsgDefLine
	return p
}

func InitEmptyEpsgDefLineContext(p *EpsgDefLineContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_epsgDefLine
}

func (*EpsgDefLineContext) IsEpsgDefLineContext() {}

func NewEpsgDefLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EpsgDefLineContext {
	var p = new(EpsgDefLineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_epsgDefLine

	return p
}

func (s *EpsgDefLineContext) GetParser() antlr.Parser { return s.parser }

func (s *EpsgDefLineContext) EpsgCode() IEpsgCodeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsgCodeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsgCodeContext)
}

func (s *EpsgDefLineContext) EQ() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserEQ, 0)
}

func (s *EpsgDefLineContext) Wkt() IWktContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWktContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWktContext)
}

func (s *EpsgDefLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpsgDefLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) EpsgDefLine() (localctx IEpsgDefLineContext) {
	localctx = NewEpsgDefLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, wktcrsv1ParserRULE_epsgDefLine)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.EpsgCode()
	}
	{
		p.SetState(97)
		p.Match(wktcrsv1ParserEQ)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Wkt()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IWktContext is an interface to support dynamic dispatch.
type IWktContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Compdcs() ICompdcsContext
	Projcs() IProjcsContext
	Geogcs() IGeogcsContext
	Vertcs() IVertcsContext
	Geoccs() IGeoccsContext
	Localcs() ILocalcsContext

	// IsWktContext differentiates from other interfaces.
	IsWktContext()
}

type WktContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWktContext() *WktContext {
	var p = new(WktContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_wkt
	return p
}

func InitEmptyWktContext(p *WktContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_wkt
}

func (*WktContext) IsWktContext() {}

func NewWktContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WktContext {
	var p = new(WktContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_wkt

	return p
}

func (s *WktContext) GetParser() antlr.Parser { return s.parser }

func (s *WktContext) Compdcs() ICompdcsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompdcsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompdcsContext)
}

func (s *WktContext) Projcs() IProjcsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjcsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjcsContext)
}

func (s *WktContext) Geogcs() IGeogcsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeogcsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeogcsContext)
}

func (s *WktContext) Vertcs() IVertcsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVertcsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVertcsContext)
}

func (s *WktContext) Geoccs() IGeoccsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeoccsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeoccsContext)
}

func (s *WktContext) Localcs() ILocalcsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalcsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalcsContext)
}

func (s *WktContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WktContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Wkt() (localctx IWktContext) {
	localctx = NewWktContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, wktcrsv1ParserRULE_wkt)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case wktcrsv1ParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Compdcs()
		}


	case wktcrsv1ParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Projcs()
		}


	case wktcrsv1ParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(102)
			p.Geogcs()
		}


	case wktcrsv1ParserT__4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)
			p.Vertcs()
		}


	case wktcrsv1ParserT__2:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(104)
			p.Geoccs()
		}


	case wktcrsv1ParserT__5:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(105)
			p.Localcs()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ICompdcsContext is an interface to support dynamic dispatch.
type ICompdcsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Vertcs() IVertcsContext
	Authority() IAuthorityContext
	RPAR() antlr.TerminalNode
	Projcs() IProjcsContext
	Geogcs() IGeogcsContext

	// IsCompdcsContext differentiates from other interfaces.
	IsCompdcsContext()
}

type CompdcsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompdcsContext() *CompdcsContext {
	var p = new(CompdcsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_compdcs
	return p
}

func InitEmptyCompdcsContext(p *CompdcsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_compdcs
}

func (*CompdcsContext) IsCompdcsContext() {}

func NewCompdcsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompdcsContext {
	var p = new(CompdcsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_compdcs

	return p
}

func (s *CompdcsContext) GetParser() antlr.Parser { return s.parser }

func (s *CompdcsContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *CompdcsContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *CompdcsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktcrsv1ParserCOMMA)
}

func (s *CompdcsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, i)
}

func (s *CompdcsContext) Vertcs() IVertcsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVertcsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVertcsContext)
}

func (s *CompdcsContext) Authority() IAuthorityContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *CompdcsContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *CompdcsContext) Projcs() IProjcsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjcsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjcsContext)
}

func (s *CompdcsContext) Geogcs() IGeogcsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeogcsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeogcsContext)
}

func (s *CompdcsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompdcsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Compdcs() (localctx ICompdcsContext) {
	localctx = NewCompdcsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, wktcrsv1ParserRULE_compdcs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(wktcrsv1ParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Name()
	}
	{
		p.SetState(111)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case wktcrsv1ParserT__1:
		{
			p.SetState(112)
			p.Projcs()
		}


	case wktcrsv1ParserT__3:
		{
			p.SetState(113)
			p.Geogcs()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(116)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Vertcs()
	}
	{
		p.SetState(118)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Authority()
	}
	{
		p.SetState(120)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IProjcsContext is an interface to support dynamic dispatch.
type IProjcsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Geogcs() IGeogcsContext
	Projection() IProjectionContext
	Unit() IUnitContext
	Authority() IAuthorityContext
	RPAR() antlr.TerminalNode
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllAxis() []IAxisContext
	Axis(i int) IAxisContext

	// IsProjcsContext differentiates from other interfaces.
	IsProjcsContext()
}

type ProjcsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjcsContext() *ProjcsContext {
	var p = new(ProjcsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_projcs
	return p
}

func InitEmptyProjcsContext(p *ProjcsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_projcs
}

func (*ProjcsContext) IsProjcsContext() {}

func NewProjcsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjcsContext {
	var p = new(ProjcsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_projcs

	return p
}

func (s *ProjcsContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjcsContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *ProjcsContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ProjcsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktcrsv1ParserCOMMA)
}

func (s *ProjcsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, i)
}

func (s *ProjcsContext) Geogcs() IGeogcsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeogcsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeogcsContext)
}

func (s *ProjcsContext) Projection() IProjectionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectionContext)
}

func (s *ProjcsContext) Unit() IUnitContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *ProjcsContext) Authority() IAuthorityContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *ProjcsContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *ProjcsContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ProjcsContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ProjcsContext) AllAxis() []IAxisContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAxisContext); ok {
			len++
		}
	}

	tst := make([]IAxisContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAxisContext); ok {
			tst[i] = t.(IAxisContext)
			i++
		}
	}

	return tst
}

func (s *ProjcsContext) Axis(i int) IAxisContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAxisContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAxisContext)
}

func (s *ProjcsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjcsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Projcs() (localctx IProjcsContext) {
	localctx = NewProjcsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, wktcrsv1ParserRULE_projcs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(wktcrsv1ParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Name()
	}
	{
		p.SetState(125)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Geogcs()
	}
	{
		p.SetState(127)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Projection()
	}
	{
		p.SetState(129)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == wktcrsv1ParserT__16 {
		{
			p.SetState(130)
			p.Parameter()
		}
		{
			p.SetState(131)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(137)
		p.Unit()
	}
	{
		p.SetState(138)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == wktcrsv1ParserT__14 {
		{
			p.SetState(139)
			p.Axis()
		}
		{
			p.SetState(140)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(147)
		p.Authority()
	}
	{
		p.SetState(148)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IGeoccsContext is an interface to support dynamic dispatch.
type IGeoccsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Datum() IDatumContext
	Primem() IPrimemContext
	Unit() IUnitContext
	Authority() IAuthorityContext
	RPAR() antlr.TerminalNode
	AllAxis() []IAxisContext
	Axis(i int) IAxisContext

	// IsGeoccsContext differentiates from other interfaces.
	IsGeoccsContext()
}

type GeoccsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeoccsContext() *GeoccsContext {
	var p = new(GeoccsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_geoccs
	return p
}

func InitEmptyGeoccsContext(p *GeoccsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_geoccs
}

func (*GeoccsContext) IsGeoccsContext() {}

func NewGeoccsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeoccsContext {
	var p = new(GeoccsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_geoccs

	return p
}

func (s *GeoccsContext) GetParser() antlr.Parser { return s.parser }

func (s *GeoccsContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *GeoccsContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *GeoccsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktcrsv1ParserCOMMA)
}

func (s *GeoccsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, i)
}

func (s *GeoccsContext) Datum() IDatumContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatumContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatumContext)
}

func (s *GeoccsContext) Primem() IPrimemContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimemContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimemContext)
}

func (s *GeoccsContext) Unit() IUnitContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *GeoccsContext) Authority() IAuthorityContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *GeoccsContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *GeoccsContext) AllAxis() []IAxisContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAxisContext); ok {
			len++
		}
	}

	tst := make([]IAxisContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAxisContext); ok {
			tst[i] = t.(IAxisContext)
			i++
		}
	}

	return tst
}

func (s *GeoccsContext) Axis(i int) IAxisContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAxisContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAxisContext)
}

func (s *GeoccsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeoccsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Geoccs() (localctx IGeoccsContext) {
	localctx = NewGeoccsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, wktcrsv1ParserRULE_geoccs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(wktcrsv1ParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Name()
	}
	{
		p.SetState(153)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Datum()
	}
	{
		p.SetState(155)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Primem()
	}
	{
		p.SetState(157)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Unit()
	}
	{
		p.SetState(159)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == wktcrsv1ParserT__14 {
		{
			p.SetState(160)
			p.Axis()
		}
		{
			p.SetState(161)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(167)
		p.Authority()
	}
	{
		p.SetState(168)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IGeogcsContext is an interface to support dynamic dispatch.
type IGeogcsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Datum() IDatumContext
	Primem() IPrimemContext
	Unit() IUnitContext
	RPAR() antlr.TerminalNode
	AllAxis() []IAxisContext
	Axis(i int) IAxisContext
	Authority() IAuthorityContext

	// IsGeogcsContext differentiates from other interfaces.
	IsGeogcsContext()
}

type GeogcsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeogcsContext() *GeogcsContext {
	var p = new(GeogcsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_geogcs
	return p
}

func InitEmptyGeogcsContext(p *GeogcsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_geogcs
}

func (*GeogcsContext) IsGeogcsContext() {}

func NewGeogcsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeogcsContext {
	var p = new(GeogcsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_geogcs

	return p
}

func (s *GeogcsContext) GetParser() antlr.Parser { return s.parser }

func (s *GeogcsContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *GeogcsContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *GeogcsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktcrsv1ParserCOMMA)
}

func (s *GeogcsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, i)
}

func (s *GeogcsContext) Datum() IDatumContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatumContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatumContext)
}

func (s *GeogcsContext) Primem() IPrimemContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimemContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimemContext)
}

func (s *GeogcsContext) Unit() IUnitContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *GeogcsContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *GeogcsContext) AllAxis() []IAxisContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAxisContext); ok {
			len++
		}
	}

	tst := make([]IAxisContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAxisContext); ok {
			tst[i] = t.(IAxisContext)
			i++
		}
	}

	return tst
}

func (s *GeogcsContext) Axis(i int) IAxisContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAxisContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAxisContext)
}

func (s *GeogcsContext) Authority() IAuthorityContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *GeogcsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeogcsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Geogcs() (localctx IGeogcsContext) {
	localctx = NewGeogcsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, wktcrsv1ParserRULE_geogcs)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(wktcrsv1ParserT__3)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(172)
		p.Name()
	}
	{
		p.SetState(173)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Datum()
	}
	{
		p.SetState(175)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Primem()
	}
	{
		p.SetState(177)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Unit()
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(179)
				p.Match(wktcrsv1ParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(180)
				p.Axis()
			}


		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == wktcrsv1ParserCOMMA {
		{
			p.SetState(186)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Authority()
		}

	}
	{
		p.SetState(190)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IVertcsContext is an interface to support dynamic dispatch.
type IVertcsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Vertdatum() IVertdatumContext
	Unit() IUnitContext
	Axis() IAxisContext
	Authority() IAuthorityContext
	RPAR() antlr.TerminalNode

	// IsVertcsContext differentiates from other interfaces.
	IsVertcsContext()
}

type VertcsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVertcsContext() *VertcsContext {
	var p = new(VertcsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_vertcs
	return p
}

func InitEmptyVertcsContext(p *VertcsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_vertcs
}

func (*VertcsContext) IsVertcsContext() {}

func NewVertcsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VertcsContext {
	var p = new(VertcsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_vertcs

	return p
}

func (s *VertcsContext) GetParser() antlr.Parser { return s.parser }

func (s *VertcsContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *VertcsContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *VertcsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktcrsv1ParserCOMMA)
}

func (s *VertcsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, i)
}

func (s *VertcsContext) Vertdatum() IVertdatumContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVertdatumContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVertdatumContext)
}

func (s *VertcsContext) Unit() IUnitContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *VertcsContext) Axis() IAxisContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAxisContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAxisContext)
}

func (s *VertcsContext) Authority() IAuthorityContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *VertcsContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *VertcsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VertcsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Vertcs() (localctx IVertcsContext) {
	localctx = NewVertcsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, wktcrsv1ParserRULE_vertcs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(wktcrsv1ParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(194)
		p.Name()
	}
	{
		p.SetState(195)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(196)
		p.Vertdatum()
	}
	{
		p.SetState(197)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(198)
		p.Unit()
	}
	{
		p.SetState(199)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(200)
		p.Axis()
	}
	{
		p.SetState(201)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Authority()
	}
	{
		p.SetState(203)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ILocalcsContext is an interface to support dynamic dispatch.
type ILocalcsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Localdatum() ILocaldatumContext
	Unit() IUnitContext
	Authority() IAuthorityContext
	RPAR() antlr.TerminalNode
	AllAxis() []IAxisContext
	Axis(i int) IAxisContext

	// IsLocalcsContext differentiates from other interfaces.
	IsLocalcsContext()
}

type LocalcsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalcsContext() *LocalcsContext {
	var p = new(LocalcsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_localcs
	return p
}

func InitEmptyLocalcsContext(p *LocalcsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_localcs
}

func (*LocalcsContext) IsLocalcsContext() {}

func NewLocalcsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalcsContext {
	var p = new(LocalcsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_localcs

	return p
}

func (s *LocalcsContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalcsContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *LocalcsContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LocalcsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktcrsv1ParserCOMMA)
}

func (s *LocalcsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, i)
}

func (s *LocalcsContext) Localdatum() ILocaldatumContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocaldatumContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocaldatumContext)
}

func (s *LocalcsContext) Unit() IUnitContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *LocalcsContext) Authority() IAuthorityContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *LocalcsContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *LocalcsContext) AllAxis() []IAxisContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAxisContext); ok {
			len++
		}
	}

	tst := make([]IAxisContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAxisContext); ok {
			tst[i] = t.(IAxisContext)
			i++
		}
	}

	return tst
}

func (s *LocalcsContext) Axis(i int) IAxisContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAxisContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAxisContext)
}

func (s *LocalcsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalcsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Localcs() (localctx ILocalcsContext) {
	localctx = NewLocalcsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, wktcrsv1ParserRULE_localcs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(wktcrsv1ParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Name()
	}
	{
		p.SetState(208)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Localdatum()
	}
	{
		p.SetState(210)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Unit()
	}
	{
		p.SetState(212)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == wktcrsv1ParserT__14 {
		{
			p.SetState(213)
			p.Axis()
		}
		{
			p.SetState(214)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(220)
		p.Authority()
	}
	{
		p.SetState(221)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IDatumContext is an interface to support dynamic dispatch.
type IDatumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Spheroid() ISpheroidContext
	RPAR() antlr.TerminalNode
	Towgs84() ITowgs84Context
	Authority() IAuthorityContext

	// IsDatumContext differentiates from other interfaces.
	IsDatumContext()
}

type DatumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatumContext() *DatumContext {
	var p = new(DatumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_datum
	return p
}

func InitEmptyDatumContext(p *DatumContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_datum
}

func (*DatumContext) IsDatumContext() {}

func NewDatumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatumContext {
	var p = new(DatumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_datum

	return p
}

func (s *DatumContext) GetParser() antlr.Parser { return s.parser }

func (s *DatumContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *DatumContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DatumContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktcrsv1ParserCOMMA)
}

func (s *DatumContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, i)
}

func (s *DatumContext) Spheroid() ISpheroidContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpheroidContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpheroidContext)
}

func (s *DatumContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *DatumContext) Towgs84() ITowgs84Context {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITowgs84Context); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITowgs84Context)
}

func (s *DatumContext) Authority() IAuthorityContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *DatumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Datum() (localctx IDatumContext) {
	localctx = NewDatumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, wktcrsv1ParserRULE_datum)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(wktcrsv1ParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Name()
	}
	{
		p.SetState(226)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(227)
		p.Spheroid()
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(228)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Towgs84()
		}

		} else if p.HasError() { // JIM
			goto errorExit
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == wktcrsv1ParserCOMMA {
		{
			p.SetState(232)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Authority()
		}

	}
	{
		p.SetState(236)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IVertdatumContext is an interface to support dynamic dispatch.
type IVertdatumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Type_() ITypeContext
	Authority() IAuthorityContext
	RPAR() antlr.TerminalNode

	// IsVertdatumContext differentiates from other interfaces.
	IsVertdatumContext()
}

type VertdatumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVertdatumContext() *VertdatumContext {
	var p = new(VertdatumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_vertdatum
	return p
}

func InitEmptyVertdatumContext(p *VertdatumContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_vertdatum
}

func (*VertdatumContext) IsVertdatumContext() {}

func NewVertdatumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VertdatumContext {
	var p = new(VertdatumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_vertdatum

	return p
}

func (s *VertdatumContext) GetParser() antlr.Parser { return s.parser }

func (s *VertdatumContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *VertdatumContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *VertdatumContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktcrsv1ParserCOMMA)
}

func (s *VertdatumContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, i)
}

func (s *VertdatumContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VertdatumContext) Authority() IAuthorityContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *VertdatumContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *VertdatumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VertdatumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Vertdatum() (localctx IVertdatumContext) {
	localctx = NewVertdatumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, wktcrsv1ParserRULE_vertdatum)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(wktcrsv1ParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(239)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(240)
		p.Name()
	}
	{
		p.SetState(241)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Type_()
	}
	{
		p.SetState(243)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Authority()
	}
	{
		p.SetState(245)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ILocaldatumContext is an interface to support dynamic dispatch.
type ILocaldatumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Type_() ITypeContext
	RPAR() antlr.TerminalNode
	Authority() IAuthorityContext

	// IsLocaldatumContext differentiates from other interfaces.
	IsLocaldatumContext()
}

type LocaldatumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocaldatumContext() *LocaldatumContext {
	var p = new(LocaldatumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_localdatum
	return p
}

func InitEmptyLocaldatumContext(p *LocaldatumContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_localdatum
}

func (*LocaldatumContext) IsLocaldatumContext() {}

func NewLocaldatumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocaldatumContext {
	var p = new(LocaldatumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_localdatum

	return p
}

func (s *LocaldatumContext) GetParser() antlr.Parser { return s.parser }

func (s *LocaldatumContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *LocaldatumContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LocaldatumContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktcrsv1ParserCOMMA)
}

func (s *LocaldatumContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, i)
}

func (s *LocaldatumContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *LocaldatumContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *LocaldatumContext) Authority() IAuthorityContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *LocaldatumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocaldatumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Localdatum() (localctx ILocaldatumContext) {
	localctx = NewLocaldatumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, wktcrsv1ParserRULE_localdatum)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(wktcrsv1ParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Name()
	}
	{
		p.SetState(250)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(251)
		p.Type_()
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == wktcrsv1ParserCOMMA {
		{
			p.SetState(252)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(253)
			p.Authority()
		}

	}
	{
		p.SetState(256)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ISpheroidContext is an interface to support dynamic dispatch.
type ISpheroidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	SemiMajorAxis() ISemiMajorAxisContext
	InverseFlattening() IInverseFlatteningContext
	RPAR() antlr.TerminalNode
	Authority() IAuthorityContext

	// IsSpheroidContext differentiates from other interfaces.
	IsSpheroidContext()
}

type SpheroidContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpheroidContext() *SpheroidContext {
	var p = new(SpheroidContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_spheroid
	return p
}

func InitEmptySpheroidContext(p *SpheroidContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_spheroid
}

func (*SpheroidContext) IsSpheroidContext() {}

func NewSpheroidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpheroidContext {
	var p = new(SpheroidContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_spheroid

	return p
}

func (s *SpheroidContext) GetParser() antlr.Parser { return s.parser }

func (s *SpheroidContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *SpheroidContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *SpheroidContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktcrsv1ParserCOMMA)
}

func (s *SpheroidContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, i)
}

func (s *SpheroidContext) SemiMajorAxis() ISemiMajorAxisContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISemiMajorAxisContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISemiMajorAxisContext)
}

func (s *SpheroidContext) InverseFlattening() IInverseFlatteningContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInverseFlatteningContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInverseFlatteningContext)
}

func (s *SpheroidContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *SpheroidContext) Authority() IAuthorityContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *SpheroidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpheroidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Spheroid() (localctx ISpheroidContext) {
	localctx = NewSpheroidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, wktcrsv1ParserRULE_spheroid)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Match(wktcrsv1ParserT__9)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(259)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Name()
	}
	{
		p.SetState(261)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(262)
		p.SemiMajorAxis()
	}
	{
		p.SetState(263)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(264)
		p.InverseFlattening()
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == wktcrsv1ParserCOMMA {
		{
			p.SetState(265)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(266)
			p.Authority()
		}

	}
	{
		p.SetState(269)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ITowgs84Context is an interface to support dynamic dispatch.
type ITowgs84Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Dx() IDxContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Dy() IDyContext
	Dz() IDzContext
	RPAR() antlr.TerminalNode
	Ex() IExContext
	Ey() IEyContext
	Ez() IEzContext
	Ppm() IPpmContext

	// IsTowgs84Context differentiates from other interfaces.
	IsTowgs84Context()
}

type Towgs84Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTowgs84Context() *Towgs84Context {
	var p = new(Towgs84Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_towgs84
	return p
}

func InitEmptyTowgs84Context(p *Towgs84Context)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_towgs84
}

func (*Towgs84Context) IsTowgs84Context() {}

func NewTowgs84Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Towgs84Context {
	var p = new(Towgs84Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_towgs84

	return p
}

func (s *Towgs84Context) GetParser() antlr.Parser { return s.parser }

func (s *Towgs84Context) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *Towgs84Context) Dx() IDxContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDxContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDxContext)
}

func (s *Towgs84Context) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktcrsv1ParserCOMMA)
}

func (s *Towgs84Context) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, i)
}

func (s *Towgs84Context) Dy() IDyContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDyContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDyContext)
}

func (s *Towgs84Context) Dz() IDzContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDzContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDzContext)
}

func (s *Towgs84Context) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *Towgs84Context) Ex() IExContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExContext)
}

func (s *Towgs84Context) Ey() IEyContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEyContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEyContext)
}

func (s *Towgs84Context) Ez() IEzContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEzContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEzContext)
}

func (s *Towgs84Context) Ppm() IPpmContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPpmContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPpmContext)
}

func (s *Towgs84Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Towgs84Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Towgs84() (localctx ITowgs84Context) {
	localctx = NewTowgs84Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, wktcrsv1ParserRULE_towgs84)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(wktcrsv1ParserT__10)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(272)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Dx()
	}
	{
		p.SetState(274)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(275)
		p.Dy()
	}
	{
		p.SetState(276)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(277)
		p.Dz()
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == wktcrsv1ParserCOMMA {
		{
			p.SetState(278)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Ex()
		}
		{
			p.SetState(280)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(281)
			p.Ey()
		}
		{
			p.SetState(282)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(283)
			p.Ez()
		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == wktcrsv1ParserCOMMA {
			{
				p.SetState(284)
				p.Match(wktcrsv1ParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(285)
				p.Ppm()
			}

		}

	}
	{
		p.SetState(290)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IAuthorityContext is an interface to support dynamic dispatch.
type IAuthorityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	AuthorityName() IAuthorityNameContext
	COMMA() antlr.TerminalNode
	Code() ICodeContext
	RPAR() antlr.TerminalNode

	// IsAuthorityContext differentiates from other interfaces.
	IsAuthorityContext()
}

type AuthorityContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuthorityContext() *AuthorityContext {
	var p = new(AuthorityContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_authority
	return p
}

func InitEmptyAuthorityContext(p *AuthorityContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_authority
}

func (*AuthorityContext) IsAuthorityContext() {}

func NewAuthorityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuthorityContext {
	var p = new(AuthorityContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_authority

	return p
}

func (s *AuthorityContext) GetParser() antlr.Parser { return s.parser }

func (s *AuthorityContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *AuthorityContext) AuthorityName() IAuthorityNameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityNameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityNameContext)
}

func (s *AuthorityContext) COMMA() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, 0)
}

func (s *AuthorityContext) Code() ICodeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeContext)
}

func (s *AuthorityContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *AuthorityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuthorityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Authority() (localctx IAuthorityContext) {
	localctx = NewAuthorityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, wktcrsv1ParserRULE_authority)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.Match(wktcrsv1ParserT__11)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(293)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(294)
		p.AuthorityName()
	}
	{
		p.SetState(295)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(296)
		p.Code()
	}
	{
		p.SetState(297)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IPrimemContext is an interface to support dynamic dispatch.
type IPrimemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	Longitude() ILongitudeContext
	RPAR() antlr.TerminalNode
	Authority() IAuthorityContext

	// IsPrimemContext differentiates from other interfaces.
	IsPrimemContext()
}

type PrimemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimemContext() *PrimemContext {
	var p = new(PrimemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_primem
	return p
}

func InitEmptyPrimemContext(p *PrimemContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_primem
}

func (*PrimemContext) IsPrimemContext() {}

func NewPrimemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimemContext {
	var p = new(PrimemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_primem

	return p
}

func (s *PrimemContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimemContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *PrimemContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *PrimemContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktcrsv1ParserCOMMA)
}

func (s *PrimemContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, i)
}

func (s *PrimemContext) Longitude() ILongitudeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILongitudeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILongitudeContext)
}

func (s *PrimemContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *PrimemContext) Authority() IAuthorityContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *PrimemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Primem() (localctx IPrimemContext) {
	localctx = NewPrimemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, wktcrsv1ParserRULE_primem)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(wktcrsv1ParserT__12)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(301)
		p.Name()
	}
	{
		p.SetState(302)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(303)
		p.Longitude()
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == wktcrsv1ParserCOMMA {
		{
			p.SetState(304)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Authority()
		}

	}
	{
		p.SetState(308)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	ConversionFactor() IConversionFactorContext
	RPAR() antlr.TerminalNode
	Authority() IAuthorityContext

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_unit
	return p
}

func InitEmptyUnitContext(p *UnitContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_unit
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *UnitContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *UnitContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktcrsv1ParserCOMMA)
}

func (s *UnitContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, i)
}

func (s *UnitContext) ConversionFactor() IConversionFactorContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConversionFactorContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConversionFactorContext)
}

func (s *UnitContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *UnitContext) Authority() IAuthorityContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Unit() (localctx IUnitContext) {
	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, wktcrsv1ParserRULE_unit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(wktcrsv1ParserT__13)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(311)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(312)
		p.Name()
	}
	{
		p.SetState(313)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(314)
		p.ConversionFactor()
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == wktcrsv1ParserCOMMA {
		{
			p.SetState(315)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Authority()
		}

	}
	{
		p.SetState(319)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IAxisContext is an interface to support dynamic dispatch.
type IAxisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	COMMA() antlr.TerminalNode
	AxisOrient() IAxisOrientContext
	RPAR() antlr.TerminalNode

	// IsAxisContext differentiates from other interfaces.
	IsAxisContext()
}

type AxisContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAxisContext() *AxisContext {
	var p = new(AxisContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_axis
	return p
}

func InitEmptyAxisContext(p *AxisContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_axis
}

func (*AxisContext) IsAxisContext() {}

func NewAxisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AxisContext {
	var p = new(AxisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_axis

	return p
}

func (s *AxisContext) GetParser() antlr.Parser { return s.parser }

func (s *AxisContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *AxisContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AxisContext) COMMA() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, 0)
}

func (s *AxisContext) AxisOrient() IAxisOrientContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAxisOrientContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAxisOrientContext)
}

func (s *AxisContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *AxisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AxisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Axis() (localctx IAxisContext) {
	localctx = NewAxisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, wktcrsv1ParserRULE_axis)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(wktcrsv1ParserT__14)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Name()
	}
	{
		p.SetState(324)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(325)
		p.AxisOrient()
	}
	{
		p.SetState(326)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IProjectionContext is an interface to support dynamic dispatch.
type IProjectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	RPAR() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	Authority() IAuthorityContext

	// IsProjectionContext differentiates from other interfaces.
	IsProjectionContext()
}

type ProjectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionContext() *ProjectionContext {
	var p = new(ProjectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_projection
	return p
}

func InitEmptyProjectionContext(p *ProjectionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_projection
}

func (*ProjectionContext) IsProjectionContext() {}

func NewProjectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionContext {
	var p = new(ProjectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_projection

	return p
}

func (s *ProjectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *ProjectionContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ProjectionContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *ProjectionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, 0)
}

func (s *ProjectionContext) Authority() IAuthorityContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorityContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorityContext)
}

func (s *ProjectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Projection() (localctx IProjectionContext) {
	localctx = NewProjectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, wktcrsv1ParserRULE_projection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(wktcrsv1ParserT__15)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(329)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(330)
		p.Name()
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == wktcrsv1ParserCOMMA {
		{
			p.SetState(331)
			p.Match(wktcrsv1ParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Authority()
		}

	}
	{
		p.SetState(335)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAR() antlr.TerminalNode
	Name() INameContext
	COMMA() antlr.TerminalNode
	Value() IValueContext
	RPAR() antlr.TerminalNode

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserLPAR, 0)
}

func (s *ParameterContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ParameterContext) COMMA() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserCOMMA, 0)
}

func (s *ParameterContext) Value() IValueContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ParameterContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserRPAR, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, wktcrsv1ParserRULE_parameter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(wktcrsv1ParserT__16)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(338)
		p.Match(wktcrsv1ParserLPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(339)
		p.Name()
	}
	{
		p.SetState(340)
		p.Match(wktcrsv1ParserCOMMA)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(341)
		p.Value()
	}
	{
		p.SetState(342)
		p.Match(wktcrsv1ParserRPAR)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IAuthorityNameContext is an interface to support dynamic dispatch.
type IAuthorityNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAuthorityNameContext differentiates from other interfaces.
	IsAuthorityNameContext()
}

type AuthorityNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuthorityNameContext() *AuthorityNameContext {
	var p = new(AuthorityNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_authorityName
	return p
}

func InitEmptyAuthorityNameContext(p *AuthorityNameContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_authorityName
}

func (*AuthorityNameContext) IsAuthorityNameContext() {}

func NewAuthorityNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuthorityNameContext {
	var p = new(AuthorityNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_authorityName

	return p
}

func (s *AuthorityNameContext) GetParser() antlr.Parser { return s.parser }
func (s *AuthorityNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuthorityNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) AuthorityName() (localctx IAuthorityNameContext) {
	localctx = NewAuthorityNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, wktcrsv1ParserRULE_authorityName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		_la = p.GetTokenStream().LA(1)

		if !(_la == wktcrsv1ParserT__17 || _la == wktcrsv1ParserT__18) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IAxisOrientContext is an interface to support dynamic dispatch.
type IAxisOrientContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsAxisOrientContext differentiates from other interfaces.
	IsAxisOrientContext()
}

type AxisOrientContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAxisOrientContext() *AxisOrientContext {
	var p = new(AxisOrientContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_axisOrient
	return p
}

func InitEmptyAxisOrientContext(p *AxisOrientContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_axisOrient
}

func (*AxisOrientContext) IsAxisOrientContext() {}

func NewAxisOrientContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AxisOrientContext {
	var p = new(AxisOrientContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_axisOrient

	return p
}

func (s *AxisOrientContext) GetParser() antlr.Parser { return s.parser }

func (s *AxisOrientContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AxisOrientContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AxisOrientContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) AxisOrient() (localctx IAxisOrientContext) {
	localctx = NewAxisOrientContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, wktcrsv1ParserRULE_axisOrient)
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case wktcrsv1ParserT__19:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(346)
			p.Match(wktcrsv1ParserT__19)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case wktcrsv1ParserT__20:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(347)
			p.Match(wktcrsv1ParserT__20)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case wktcrsv1ParserT__21:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(348)
			p.Match(wktcrsv1ParserT__21)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case wktcrsv1ParserT__22:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(349)
			p.Match(wktcrsv1ParserT__22)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case wktcrsv1ParserT__23:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(350)
			p.Match(wktcrsv1ParserT__23)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case wktcrsv1ParserT__24:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(351)
			p.Match(wktcrsv1ParserT__24)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case wktcrsv1ParserT__25:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(352)
			p.Match(wktcrsv1ParserT__25)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case wktcrsv1ParserT__26:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(353)
			p.Match(wktcrsv1ParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case wktcrsv1ParserT__27:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(354)
			p.Match(wktcrsv1ParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case wktcrsv1ParserT__28:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(355)
			p.Match(wktcrsv1ParserT__28)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case wktcrsv1ParserT__29:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(356)
			p.Match(wktcrsv1ParserT__29)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case wktcrsv1ParserTEXT:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(357)
			p.Name()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IEpsgCodeContext is an interface to support dynamic dispatch.
type IEpsgCodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PKEY() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsEpsgCodeContext differentiates from other interfaces.
	IsEpsgCodeContext()
}

type EpsgCodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEpsgCodeContext() *EpsgCodeContext {
	var p = new(EpsgCodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_epsgCode
	return p
}

func InitEmptyEpsgCodeContext(p *EpsgCodeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_epsgCode
}

func (*EpsgCodeContext) IsEpsgCodeContext() {}

func NewEpsgCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EpsgCodeContext {
	var p = new(EpsgCodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_epsgCode

	return p
}

func (s *EpsgCodeContext) GetParser() antlr.Parser { return s.parser }

func (s *EpsgCodeContext) PKEY() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserPKEY, 0)
}

func (s *EpsgCodeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *EpsgCodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpsgCodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) EpsgCode() (localctx IEpsgCodeContext) {
	localctx = NewEpsgCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, wktcrsv1ParserRULE_epsgCode)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		_la = p.GetTokenStream().LA(1)

		if !(_la == wktcrsv1ParserNUMBER || _la == wktcrsv1ParserPKEY) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) TEXT() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserTEXT, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, wktcrsv1ParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.Match(wktcrsv1ParserTEXT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, wktcrsv1ParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, wktcrsv1ParserRULE_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ISemiMajorAxisContext is an interface to support dynamic dispatch.
type ISemiMajorAxisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsSemiMajorAxisContext differentiates from other interfaces.
	IsSemiMajorAxisContext()
}

type SemiMajorAxisContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySemiMajorAxisContext() *SemiMajorAxisContext {
	var p = new(SemiMajorAxisContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_semiMajorAxis
	return p
}

func InitEmptySemiMajorAxisContext(p *SemiMajorAxisContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_semiMajorAxis
}

func (*SemiMajorAxisContext) IsSemiMajorAxisContext() {}

func NewSemiMajorAxisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SemiMajorAxisContext {
	var p = new(SemiMajorAxisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_semiMajorAxis

	return p
}

func (s *SemiMajorAxisContext) GetParser() antlr.Parser { return s.parser }

func (s *SemiMajorAxisContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *SemiMajorAxisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SemiMajorAxisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) SemiMajorAxis() (localctx ISemiMajorAxisContext) {
	localctx = NewSemiMajorAxisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, wktcrsv1ParserRULE_semiMajorAxis)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IInverseFlatteningContext is an interface to support dynamic dispatch.
type IInverseFlatteningContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsInverseFlatteningContext differentiates from other interfaces.
	IsInverseFlatteningContext()
}

type InverseFlatteningContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInverseFlatteningContext() *InverseFlatteningContext {
	var p = new(InverseFlatteningContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_inverseFlattening
	return p
}

func InitEmptyInverseFlatteningContext(p *InverseFlatteningContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_inverseFlattening
}

func (*InverseFlatteningContext) IsInverseFlatteningContext() {}

func NewInverseFlatteningContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InverseFlatteningContext {
	var p = new(InverseFlatteningContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_inverseFlattening

	return p
}

func (s *InverseFlatteningContext) GetParser() antlr.Parser { return s.parser }

func (s *InverseFlatteningContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *InverseFlatteningContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InverseFlatteningContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) InverseFlattening() (localctx IInverseFlatteningContext) {
	localctx = NewInverseFlatteningContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, wktcrsv1ParserRULE_inverseFlattening)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IDxContext is an interface to support dynamic dispatch.
type IDxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsDxContext differentiates from other interfaces.
	IsDxContext()
}

type DxContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDxContext() *DxContext {
	var p = new(DxContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_dx
	return p
}

func InitEmptyDxContext(p *DxContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_dx
}

func (*DxContext) IsDxContext() {}

func NewDxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DxContext {
	var p = new(DxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_dx

	return p
}

func (s *DxContext) GetParser() antlr.Parser { return s.parser }

func (s *DxContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *DxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Dx() (localctx IDxContext) {
	localctx = NewDxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, wktcrsv1ParserRULE_dx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IDyContext is an interface to support dynamic dispatch.
type IDyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsDyContext differentiates from other interfaces.
	IsDyContext()
}

type DyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDyContext() *DyContext {
	var p = new(DyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_dy
	return p
}

func InitEmptyDyContext(p *DyContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_dy
}

func (*DyContext) IsDyContext() {}

func NewDyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DyContext {
	var p = new(DyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_dy

	return p
}

func (s *DyContext) GetParser() antlr.Parser { return s.parser }

func (s *DyContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *DyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Dy() (localctx IDyContext) {
	localctx = NewDyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, wktcrsv1ParserRULE_dy)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IDzContext is an interface to support dynamic dispatch.
type IDzContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsDzContext differentiates from other interfaces.
	IsDzContext()
}

type DzContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDzContext() *DzContext {
	var p = new(DzContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_dz
	return p
}

func InitEmptyDzContext(p *DzContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_dz
}

func (*DzContext) IsDzContext() {}

func NewDzContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DzContext {
	var p = new(DzContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_dz

	return p
}

func (s *DzContext) GetParser() antlr.Parser { return s.parser }

func (s *DzContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *DzContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DzContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Dz() (localctx IDzContext) {
	localctx = NewDzContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, wktcrsv1ParserRULE_dz)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IExContext is an interface to support dynamic dispatch.
type IExContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsExContext differentiates from other interfaces.
	IsExContext()
}

type ExContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExContext() *ExContext {
	var p = new(ExContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_ex
	return p
}

func InitEmptyExContext(p *ExContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_ex
}

func (*ExContext) IsExContext() {}

func NewExContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExContext {
	var p = new(ExContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_ex

	return p
}

func (s *ExContext) GetParser() antlr.Parser { return s.parser }

func (s *ExContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *ExContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Ex() (localctx IExContext) {
	localctx = NewExContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, wktcrsv1ParserRULE_ex)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IEyContext is an interface to support dynamic dispatch.
type IEyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsEyContext differentiates from other interfaces.
	IsEyContext()
}

type EyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEyContext() *EyContext {
	var p = new(EyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_ey
	return p
}

func InitEmptyEyContext(p *EyContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_ey
}

func (*EyContext) IsEyContext() {}

func NewEyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EyContext {
	var p = new(EyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_ey

	return p
}

func (s *EyContext) GetParser() antlr.Parser { return s.parser }

func (s *EyContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *EyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Ey() (localctx IEyContext) {
	localctx = NewEyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, wktcrsv1ParserRULE_ey)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IEzContext is an interface to support dynamic dispatch.
type IEzContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsEzContext differentiates from other interfaces.
	IsEzContext()
}

type EzContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEzContext() *EzContext {
	var p = new(EzContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_ez
	return p
}

func InitEmptyEzContext(p *EzContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_ez
}

func (*EzContext) IsEzContext() {}

func NewEzContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EzContext {
	var p = new(EzContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_ez

	return p
}

func (s *EzContext) GetParser() antlr.Parser { return s.parser }

func (s *EzContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *EzContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EzContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Ez() (localctx IEzContext) {
	localctx = NewEzContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, wktcrsv1ParserRULE_ez)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IPpmContext is an interface to support dynamic dispatch.
type IPpmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsPpmContext differentiates from other interfaces.
	IsPpmContext()
}

type PpmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPpmContext() *PpmContext {
	var p = new(PpmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_ppm
	return p
}

func InitEmptyPpmContext(p *PpmContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_ppm
}

func (*PpmContext) IsPpmContext() {}

func NewPpmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PpmContext {
	var p = new(PpmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_ppm

	return p
}

func (s *PpmContext) GetParser() antlr.Parser { return s.parser }

func (s *PpmContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *PpmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PpmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Ppm() (localctx IPpmContext) {
	localctx = NewPpmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, wktcrsv1ParserRULE_ppm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(384)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ICodeContext is an interface to support dynamic dispatch.
type ICodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT() antlr.TerminalNode

	// IsCodeContext differentiates from other interfaces.
	IsCodeContext()
}

type CodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeContext() *CodeContext {
	var p = new(CodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_code
	return p
}

func InitEmptyCodeContext(p *CodeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_code
}

func (*CodeContext) IsCodeContext() {}

func NewCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeContext {
	var p = new(CodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_code

	return p
}

func (s *CodeContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeContext) TEXT() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserTEXT, 0)
}

func (s *CodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Code() (localctx ICodeContext) {
	localctx = NewCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, wktcrsv1ParserRULE_code)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.Match(wktcrsv1ParserTEXT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ILongitudeContext is an interface to support dynamic dispatch.
type ILongitudeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsLongitudeContext differentiates from other interfaces.
	IsLongitudeContext()
}

type LongitudeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLongitudeContext() *LongitudeContext {
	var p = new(LongitudeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_longitude
	return p
}

func InitEmptyLongitudeContext(p *LongitudeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_longitude
}

func (*LongitudeContext) IsLongitudeContext() {}

func NewLongitudeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LongitudeContext {
	var p = new(LongitudeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_longitude

	return p
}

func (s *LongitudeContext) GetParser() antlr.Parser { return s.parser }

func (s *LongitudeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *LongitudeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LongitudeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Longitude() (localctx ILongitudeContext) {
	localctx = NewLongitudeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, wktcrsv1ParserRULE_longitude)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(388)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IConversionFactorContext is an interface to support dynamic dispatch.
type IConversionFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsConversionFactorContext differentiates from other interfaces.
	IsConversionFactorContext()
}

type ConversionFactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConversionFactorContext() *ConversionFactorContext {
	var p = new(ConversionFactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_conversionFactor
	return p
}

func InitEmptyConversionFactorContext(p *ConversionFactorContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_conversionFactor
}

func (*ConversionFactorContext) IsConversionFactorContext() {}

func NewConversionFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConversionFactorContext {
	var p = new(ConversionFactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_conversionFactor

	return p
}

func (s *ConversionFactorContext) GetParser() antlr.Parser { return s.parser }

func (s *ConversionFactorContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *ConversionFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConversionFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) ConversionFactor() (localctx IConversionFactorContext) {
	localctx = NewConversionFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, wktcrsv1ParserRULE_conversionFactor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = wktcrsv1ParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktcrsv1ParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(wktcrsv1ParserNUMBER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




func (p *wktcrsv1Parser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, wktcrsv1ParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		p.Match(wktcrsv1ParserNUMBER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


