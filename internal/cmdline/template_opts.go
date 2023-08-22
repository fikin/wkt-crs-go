package cmdline

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/creachadair/goflags/enumflag"
	wktcrs "github.com/fikin/wkt-crs-go"
)

type configTemplateOpt struct {
	l             *log.Logger
	packageName   string
	templateFiles string
	format        *enumflag.Value
	inFile        *ioInputValue
	outFile       *ioOutputValue
}

func parseTemplateFlags(log *log.Logger, args []string) (*configTemplateOpt, func(), error) {
	c := &configTemplateOpt{
		l:       log,
		inFile:  newIoInputVar(),
		outFile: newIoOutputVar(),
		format:  enumflag.New("json", "index", "wgs84", "ast"),
	}
	fl := flag.NewFlagSet("template", flag.ContinueOnError)
	fl.StringVar(&c.packageName, "package", "", "variable \"package\" passed to the template")
	fl.Var(c.inFile, "in", "path to epsg.properties file or \"-\" for stdin")
	fl.Var(c.outFile, "out", "path to output file or \"-\" for stdout")
	fl.StringVar(&c.templateFiles, "template", "", "path and file (pattern) to go-lang template files i.e. dir/prefix\\*.tmpl")
	fl.Var(c.format, "format", "convert input to json, index or wgs84 data and bind ir to \"repo\" template variable. \"ast\" option is antlr root node as-is.")
	if err := fl.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil, func() {}, err
		}
		return nil, fl.Usage, err
	}
	if c.templateFiles == "" {
		return nil, fl.Usage, fmt.Errorf("define -templates option")
	}

	return c, fl.Usage, nil
}

func handleTemplate(cfg *configTemplateOpt) error {
	fNames, err := parseFS(os.DirFS(path.Dir(cfg.templateFiles)), []string{path.Base(cfg.templateFiles)})
	if err != nil {
		return err
	}
	tmpl, err := template.New(fNames[0]).Funcs(getFuncsMap()).ParseFS(os.DirFS(path.Dir(cfg.templateFiles)), path.Base(cfg.templateFiles))
	if err != nil {
		return err
	}
	is := antlr.NewIoStream(cfg.inFile.file)
	parser := wktcrs.NewWktcrsv1Parser(is)
	repo, err := getTransformedAST(cfg.format.Key(), parser.PropsFile())
	if err != nil {
		return err
	}
	return tmpl.Execute(cfg.outFile.file, map[string]interface{}{
		"packageName": cfg.packageName,
		"repo":        repo,
		"inFileName":  cfg.inFile,
	})
}

func getFuncsMap() template.FuncMap {
	return template.FuncMap{
		"float64": func(val any) any {
			switch n := val.(type) {
			case int:
				return float64(n)
			case float32:
				return float64(n)
			default:
				return val
			}
		},
		"nvl": func(val, def any) any {
			if val == nil {
				return def
			}
			return val
		},
		"cammelCase": func(elems ...string) string {
			arr := make([]string, len(elems))
			for i, s := range elems {
				// nolint:staticcheck
				arr[i] = strings.Title(strings.ToLower(s))
			}
			return strings.Join(arr, "")
		},
		"arr": func(elems ...any) []any {
			return elems
		},
		"strVal": func(str *string) string {
			return *str
		},
		"formatNumber": func(value float64) string {
			return strconv.FormatFloat(value, 'f', -1, 64)
		},
		"formatNumberPtr": func(value *float64) string {
			if value == nil {
				return "0"
			}
			return strconv.FormatFloat(*value, 'f', -1, 64)
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
	}
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
