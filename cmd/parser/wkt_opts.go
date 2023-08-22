package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"

	"github.com/antlr4-go/antlr/v4"
	"github.com/creachadair/goflags/enumflag"
	wktcrs "github.com/fikin/wkt-crs-go"
)

type configWktOpt struct {
	l       *log.Logger
	inFile  *ioInputValue
	outFile *ioOutputValue
	format  *enumflag.Value
}

func parseWktFlags(log *log.Logger, args []string) (*configWktOpt, func(), error) {
	c := &configWktOpt{
		l:       log,
		format:  enumflag.New("slim", "pretty", "json"),
		inFile:  newIoInputVar(),
		outFile: newIoOutputVar(),
	}
	fl := flag.NewFlagSet("wkt", flag.ContinueOnError)
	fl.Var(c.format, "format", "output WKT as slim, pretty or json")
	fl.Var(c.inFile, "in", "input file with WKT or \"-\" for stdin")
	fl.Var(c.outFile, "out", "output file or \"-\" for stdout")
	if err := fl.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil, func() {}, err
		}
		return nil, fl.Usage, err
	}
	return c, fl.Usage, nil
}

func handlerWkt(cfg *configWktOpt) error {
	scanner := bufio.NewScanner(cfg.inFile.file)
	for scanner.Scan() {
		line := scanner.Text()
		is := antlr.NewInputStream(line)
		tree := wktcrs.NewWktcrsv1Parser(is).Wkt()
		switch cfg.format.Key() {
		case "pretty":
			return addNL(wktcrs.NodeToPrettyText(tree, "", "  ", cfg.outFile), cfg.outFile)
		case "json":
			obj, err := wktcrs.NodeToMap(tree)
			if err != nil {
				return err
			}
			return addNL(toPrettyJSON(cfg.outFile.file, obj), cfg.outFile)
		default:
			return addNL(wktcrs.NodeToText(tree, cfg.outFile), cfg.outFile)
		}
	}
	return nil
}

func addNL(err error, out io.StringWriter) error {
	if err != nil {
		return err
	}
	_, err = out.WriteString("\n")
	return err
}
