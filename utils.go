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

// ParsePropsFileAST is parsing input of epsg.properties file and
// returns AST root node.
func ParsePropsFileAST(is antlr.CharStream) wktcrsv1.IPropsFileContext {
	lexer := wktcrsv1.Newwktcrsv1Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := wktcrsv1.Newwktcrsv1Parser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	return parser.PropsFile()
}
