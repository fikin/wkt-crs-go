// Package main is providing with cmdline tool to process WKT or epsg.properties
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/antlr4-go/antlr/v4"
	wktcrs "github.com/fikin/wkt-crs-go"
)

func main() {
	log := log.Default()

	fn := parseMainOpts(log, os.Args[1:], handlerWkt, handleTemplate)

	// nolint:staticcheck
	if usage, err := fn(); err != nil {
		log.Fatalf("%#v", err)
		usage()
		os.Exit(1)
	}
}

func handleTemplate(cfg *configPropsOpt) error {
	fNames, err := parseFS(os.DirFS(path.Dir(cfg.templateFiles)), []string{path.Base(cfg.templateFiles)})
	if err != nil {
		return err
	}
	tmpl, err := template.New(fNames[0]).Funcs(template.FuncMap{
		"strVal": func(str *string) string {
			return *str
		},
		"formatNumber": func(value float64) string {
			return fmt.Sprintf("%.2f", value)
		},
		"formatNumberPtr": func(value *float64) string {
			if value == nil {
				return "0"
			}
			return fmt.Sprintf("%.2f", *value)
		},
		"toid": func(str string) string {
			return strings.ReplaceAll(str, " ", "")
		},
		"totext": func(n antlr.Tree) (string, error) {
			buf := strings.Builder{}
			if err := wktcrs.NodeToText(n, &buf); err != nil {
				return "", err
			}
			return buf.String(), nil
		},
		"prettytext": func(n antlr.Tree) (string, error) {
			buf := strings.Builder{}
			if err := wktcrs.NodeToPrettyText(n, "", "  ", &buf); err != nil {
				return "", err
			}
			return buf.String(), nil
		},
		"lookupProjectionMethod": func(code string) []string {
			switch code {
			case "1024": // "Popular Visualisation Pseudo Mercator",
				return []string{"WebMercator"}
			case "1027": // "Lambert Azimuthal Equal Area",
				return []string{"LambertAzimuthalEqualArea", "LatitudeOfCenter", "LongitudeOfCenter", "FalseEasting", "FalseNorthing"}
			// case "1028": // "Equidistant Cylindrical",
			// case "1029": // "Equidistant Cylindrical Spherical",
			// case "1078": // "Equal Earth",
			// case "9801": // "Lambert Conformal Conic 1sp",
			case "9802": // "Lambert Conformal Conic 2sp",
				return []string{"LambertConformalConic2SP", "CentralMeridian", "LatitudeOfOrigin", "StandardParallel1", "StandardParallel2", "FalseEasting", "FalseNorthing"}
			// case "9803": // "Lambert Conformal Conic 2sp Belgium",
			// case "9804": // "Mercator 1sp",
			// case "9805": // "Mercator 2sp",
			// case "9806": // "Cassini Soldner",
			case "9807": // "Transverse Mercator",
				return []string{"TransverseMercator", "CentralMeridian", "LatitudeOfOrigin", "ScaleFactor", "FalseEasting", "FalseNorthing"}
			// case "9808": // "Transverse Mercator South Orientated",
			// case "9809": // "Oblique Stereographic",
			// case "9810": // "Polar Stereographic",
			// case "9811": // "New Zealand Map Grid",
			// case "9812": // "Hotine Oblique Mercator",
			// case "9815": // "Oblique Mercator",
			// case "9818": // "Polyconic",
			// case "9819": // "Krovak",
			// case "9820": // "Lambert Azimuthal Equal Area",
			case "9822": // "Albers Conic Equal Area",
				return []string{"AlbersEqualAreaConic", "StandardParallel1", "StandardParallel2", "LatitudeOfOrigin", "CentralMeridian", "FalseEasting", "FalseNorthing"}
			// case "9823": // "Equidistant Cylindrical Spherical",
			// case "9829": // "Polar Stereographic Variant B",
			// case "9834": // "Lambert Cylindrical Equal Area Spherical",
			// case "9841": // "Mercator 1sp",
			// case "9842": // "Equidistant Cylindrical",
			// case "Stereographic": // "Stereographic"
			default:
				return nil
			}
		},
	}).ParseFS(os.DirFS(path.Dir(cfg.templateFiles)), path.Base(cfg.templateFiles))
	if err != nil {
		return err
	}
	is, err := antlr.NewFileStream(cfg.propsFile)
	if err != nil {
		return err
	}
	os, err := os.Create(cfg.outFile)
	if err != nil {
		return err
	}
	parser := wktcrs.NewWktcrsv1Parser(is)
	indx, err := wktcrs.AuthoritiesIndexPropFile(parser.PropsFile())
	if err != nil {
		return err
	}
	return tmpl.Execute(os, map[string]interface{}{
		"packageName": cfg.packageName,
		"repo":        indx,
		"inFileName":  cfg.propsFile,
	})
}

type configPropsOpt struct {
	log           *log.Logger
	packageName   string
	version       string
	propsFile     string
	templateFiles string
	outFile       string
}

type configWktOpt struct {
	log    *log.Logger
	pretty bool
}

func parseMainOpts(
	l *log.Logger,
	args []string,
	onWkt func(*configWktOpt, io.Reader, io.StringWriter) error,
	onProps func(*configPropsOpt) error,
) func() (func(), error) {
	ffl := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	if err := ffl.Parse(args); err != nil {
		return func() (func(), error) {
			return ffl.Usage, err
		}
	}
	cmd, args2 := func() (string, []string) {
		arr := ffl.Args()
		if len(arr) == 0 {
			return "err", nil
		}
		return arr[0], arr[1:]
	}()
	switch cmd {
	case "wkt":
		cfg, fl, err := parseWktFlags(l, args2)
		return func() (func(), error) {
			if err != nil {
				return fl, err
			}
			return fl, onWkt(cfg, os.Stdin, os.Stdout)
		}
	case "props":
		cfg, fl, err := parsePropsFlags(l, args2)
		return func() (func(), error) {
			if err != nil {
				return fl, err
			}
			return fl, onProps(cfg)
		}
	default:
		return func() (func(), error) {
			return flag.Usage, fmt.Errorf("pls choose \"wkt\" or \"props\" option")
		}
	}
}

func handlerWkt(cfg *configWktOpt, in io.Reader, out io.StringWriter) error {
	is := antlr.NewIoStream(bufio.NewReader(in))
	tree := wktcrs.NewWktcrsv1Parser(is).Wkt()
	switch {
	case cfg.pretty:
		return wktcrs.NodeToPrettyText(tree, "", "  ", out)
	default:
		return wktcrs.NodeToText(tree, out)
	}
}

func parseWktFlags(log *log.Logger, args []string) (*configWktOpt, func(), error) {
	c := &configWktOpt{
		log: log,
	}
	fl := flag.NewFlagSet("wkt", flag.ContinueOnError)
	fl.BoolVar(&c.pretty, "pretty", false, "print pretty")
	if err := fl.Parse(args); err != nil {
		return nil, fl.Usage, err
	}
	return c, fl.Usage, nil
}

func parsePropsFlags(log *log.Logger, args []string) (*configPropsOpt, func(), error) {
	c := &configPropsOpt{
		log: log,
	}
	fl := flag.NewFlagSet("props", flag.ContinueOnError)
	fl.StringVar(&c.version, "version", "v1", "v1 or v2")
	fl.StringVar(&c.packageName, "package", "", "package name of the generated code")
	fl.StringVar(&c.propsFile, "in", "", "path to epsg.properties file")
	fl.StringVar(&c.templateFiles, "templates", "", "path and file pattern to go-lang template files")
	fl.StringVar(&c.outFile, "out", "", "path to file to save the template output")
	if err := fl.Parse(args); err != nil {
		return nil, fl.Usage, err
	}
	switch {
	case c.packageName == "":
		return nil, fl.Usage, fmt.Errorf("define -package option")
	case c.propsFile == "":
		return nil, fl.Usage, fmt.Errorf("define -in option")
	case c.templateFiles == "":
		return nil, fl.Usage, fmt.Errorf("define -templates option")
	case c.outFile == "":
		return nil, fl.Usage, fmt.Errorf("define -out option")
	}

	return c, fl.Usage, nil
}

func parseFS(fsys fs.FS, patterns []string) ([]string, error) {
	var filenames []string
	for _, pattern := range patterns {
		list, err := fs.Glob(fsys, pattern)
		if err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, fmt.Errorf("template: pattern matches no files: %#q", pattern)
		}
		filenames = append(filenames, list...)
	}
	return filenames, nil
}
