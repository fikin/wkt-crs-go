// Package wktcrs is providing tools on top of WKT CRS parsers.
package wktcrs

import (
	"fmt"
	"io"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fikin/wkt-crs-go/wktcrsv1"
)

type (
	newLineWriter func(io.StringWriter, string)
	visitorFn     func(antlr.Tree) error
)

func writeString(out io.StringWriter, msg string, args ...interface{}) error {
	_, err := out.WriteString(fmt.Sprintf(msg, args...))
	return err
}

func writeOnNewLine() newLineWriter {
	firstTime := true
	return func(out io.StringWriter, str string) {
		switch firstTime {
		case true:
			firstTime = false
		default:
			_, _ = out.WriteString("\n")
		}
		_, _ = out.WriteString(str)
	}
}

func visitChildren(node antlr.Tree, visitor visitorFn) error {
	for _, c := range node.GetChildren() {
		if err := visitor(c); err != nil {
			return err
		}
	}
	return nil
}

// NewWktcrsv1Parser is returning parser for given input.
func NewWktcrsv1Parser(is antlr.CharStream) *wktcrsv1.Wktcrsv1Parser {
	lexer := wktcrsv1.Newwktcrsv1Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := wktcrsv1.NewWktcrsv1Parser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	return parser
}
