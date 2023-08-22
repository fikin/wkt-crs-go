// Package parser is providing with cmdline tool to process WKT or epsg.properties
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	l := log.Default()

	fn := parseMainOpts(l, os.Args[1:], handlerWkt, handleProps, handleTemplate)

	// nolint:staticcheck
	if usage, err := fn(); err != nil {
		l.Printf("[ERR] %v\n", err)
		usage()
		os.Exit(1)
	}
}

func parseMainOpts(
	l *log.Logger,
	args []string,
	onWkt func(*configWktOpt) error,
	onProps func(*configPropsOpt) error,
	onTemplate func(*configTemplateOpt) error,
) func() (func(), error) {
	ffl := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	ffl.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: [wkt ... | props ... | -h]\n")
		fmt.Fprintf(os.Stderr, "  wkt\n")
		fmt.Fprintf(os.Stderr, "         reads WKT from stdin and prints it to stdout\n")
		fmt.Fprintf(os.Stderr, "  props\n")
		fmt.Fprintf(os.Stderr, "         reads epsg.properties-like file and transforms it\n")
		fmt.Fprintf(os.Stderr, "  template\n")
		fmt.Fprintf(os.Stderr, "         reads epsg.properties-like file and applies templates to it\n")
	}
	if err := ffl.Parse(args); err != nil {
		return func() (func(), error) {
			if errors.Is(err, flag.ErrHelp) {
				return func() {}, err
			}
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
			return fl, onWkt(cfg)
		}
	case "template":
		cfg, fl, err := parseTemplateFlags(l, args2)
		return func() (func(), error) {
			if err != nil {
				return fl, err
			}
			return fl, onTemplate(cfg)
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
			return ffl.Usage, fmt.Errorf("pls choose \"wkt\" or \"props\" option")
		}
	}
}
