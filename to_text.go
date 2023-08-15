package wktcrs

import (
	"io"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fikin/wkt-crs-go/wktcrsv1"
)

// NodeToText is printing given node as WKT text, no whitespace used.
func NodeToText(node antlr.Tree, out io.StringWriter) error {
	return nodeToText2(node, out, writeOnNewLine())
}

func nodeToText2(node antlr.Tree, out io.StringWriter, nl newLineWriter) error {
	switch n := node.(type) {
	case wktcrsv1.IPropRowContext:
		nl(out, "")
		return visitChildren(n, func(t antlr.Tree) error { return nodeToText2(t, out, nl) })
	case antlr.ErrorNode:
		return writeString(out, ">>%s<<", n.(antlr.ParseTree).GetText())
	case antlr.TerminalNode:
		switch n.GetSymbol().GetTokenType() {
		case antlr.TokenEOF:
			return nil
		default:
			return writeString(out, n.(antlr.ParseTree).GetText())
		}
	default:
		return visitChildren(n, func(t antlr.Tree) error { return nodeToText2(t, out, nl) })
	}
}
