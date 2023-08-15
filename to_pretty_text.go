package wktcrs

import (
	"io"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fikin/wkt-crs-go/wktcrsv1"
)

// NodeToPrettyText is returning node as pretty printed.
func NodeToPrettyText(node antlr.Tree, prefix, indent string, out io.StringWriter) error {
	return nodeToPrettyText2(node, out, writeOnNewLine(), prefix, indent)
}

func nodeToPrettyText2(node antlr.Tree, out io.StringWriter, nl newLineWriter, prefix, indent string) error {
	switch n := node.(type) {
	case wktcrsv1.IAuthorityContext,
		wktcrsv1.ICompdcsContext,
		wktcrsv1.IGeoccsContext,
		wktcrsv1.ILocalcsContext,
		wktcrsv1.ILocaldatumContext,
		wktcrsv1.IPrimemContext,
		wktcrsv1.IParameterContext,
		wktcrsv1.IProjcsContext,
		wktcrsv1.IProjectionContext,
		wktcrsv1.ISpheroidContext,
		wktcrsv1.ITowgs84Context,
		wktcrsv1.IGeogcsContext,
		wktcrsv1.IAxisContext,
		wktcrsv1.IDatumContext,
		wktcrsv1.IUnitContext,
		wktcrsv1.IVertcsContext,
		wktcrsv1.IVertdatumContext,
		wktcrsv1.IPropRowContext:
		nl(out, prefix)
		return visitChildren(n, func(t antlr.Tree) error { return nodeToPrettyText2(t, out, nl, prefix+indent, indent) })
	case antlr.ErrorNode:
		return writeString(out, ">>%s<<", n.(antlr.ParseTree).GetText())
	case antlr.TerminalNode:
		switch n.GetSymbol().GetTokenType() {
		case antlr.TokenEOF:
			return nil
		case wktcrsv1.Wktcrsv1LexerCOMMA:
			return writeString(out, "%s ", n.(antlr.ParseTree).GetText())
		default:
			return writeString(out, n.(antlr.ParseTree).GetText())
		}
	default:
		return visitChildren(n, func(t antlr.Tree) error { return nodeToPrettyText2(t, out, nl, prefix, indent) })
	}
}
