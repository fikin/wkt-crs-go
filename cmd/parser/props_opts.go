package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/antlr4-go/antlr/v4"
	"github.com/creachadair/goflags/enumflag"
	wktcrs "github.com/fikin/wkt-crs-go"
	"github.com/fikin/wkt-crs-go/wktcrsv1"
)

type configPropsOpt struct {
	l       *log.Logger
	format  *enumflag.Value
	inFile  *ioInputValue
	outFile *ioOutputValue
}

func parsePropsFlags(log *log.Logger, args []string) (*configPropsOpt, func(), error) {
	c := &configPropsOpt{
		l:       log,
		inFile:  newIoInputVar(),
		outFile: newIoOutputVar(),
		format:  enumflag.New("json", "index", "wgs84", "pretty", "slim"),
	}
	fl := flag.NewFlagSet("props", flag.ContinueOnError)
	fl.Var(c.inFile, "in", "path to epsg.properties file or \"-\" for stdin")
	fl.Var(c.outFile, "out", "path to output file or \"-\" for stdout")
	fl.Var(c.format, "format", "output all CRS as \"json\", \"pretty\" or \"slim\" text, \"index\" all CRS or only the ones with to-\"wgs84\" definitions and print them as json")
	if err := fl.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil, func() {}, err
		}
		return nil, fl.Usage, err
	}

	return c, fl.Usage, nil
}

func handleProps(cfg *configPropsOpt) error {
	is := antlr.NewIoStream(cfg.inFile.file)
	parser := wktcrs.NewWktcrsv1Parser(is)
	switch cfg.format.Key() {
	case "json", "index", "wgs84":
		objs, err := getTransformedAST(cfg.format.Key(), parser.PropsFile())
		if err != nil {
			return err
		}
		return toPrettyJSON(cfg.outFile.file, objs)
	case "pretty":
		return wktcrs.NodeToPrettyText(parser.PropsFile(), "", "  ", cfg.outFile)
	default:
		return wktcrs.NodeToText(parser.PropsFile(), cfg.outFile)
	}
}

func toPrettyJSON(out io.Writer, data any) error {
	e := json.NewEncoder(out)
	e.SetIndent("", "  ")
	return e.Encode(data)
}

func getTransformedAST(format string, node wktcrsv1.IPropsFileContext) (any, error) {
	switch format {
	case "json":
		return wktcrs.PropsFileNodeToArr(node)
	case "index":
		return wktcrs.AuthoritiesIndexPropFile(node)
	case "wgs84":
		return wktcrs.Wgs84PropFile(node)
	default:
		return nil, fmt.Errorf("unsupported format options : %s", format)
	}
}
