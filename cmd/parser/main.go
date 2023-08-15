package main

import (
	"flag"
	"fmt"
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
	cfg := parseFlags()

	if err := handleTemplate(cfg); err != nil {
		cfg.log.Fatalf("%#v", err)
		os.Exit(1)
	}
}

func handleTemplate(cfg config) error {
	fNames, err := parseFS(os.DirFS(path.Dir(cfg.templateFiles)), []string{path.Base(cfg.templateFiles)})
	if err != nil {
		return err
	}
	tmpl, err := template.New(fNames[0]).Funcs(template.FuncMap{
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
	propsFileNode := wktcrs.ParsePropsFileAST(is)
	return tmpl.Execute(os, map[string]interface{}{
		"packageName": cfg.packageName,
		"propsNode":   propsFileNode,
		"inFileName":  cfg.propsFile,
	})
}

type config struct {
	log           *log.Logger
	packageName   string
	version       string
	propsFile     string
	templateFiles string
	outFile       string
}

func parseFlags() config {
	c := config{
		log: log.Default(),
	}
	flag.StringVar(&c.version, "version", "v1", "v1 or v2")
	flag.StringVar(&c.packageName, "package", "", "package name of the generated code")
	flag.StringVar(&c.propsFile, "properties", "", "path to epsg.properties file")
	flag.StringVar(&c.templateFiles, "templates", "", "path and file pattern to go-lang template files")
	flag.StringVar(&c.outFile, "out", "", "path to file to save the template output")
	flag.Parse()
	return c
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
