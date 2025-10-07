package ast

import (
	"fmt"
	"strings"
)

type AstPrinter struct{}

func (ap *AstPrinter) VisitBinaryExpr(e Binary[string]) string {
	return ap.parenthesize(e.Operator.Lexeme, e.Left, e.Right)
}

func (ap *AstPrinter) VisitUnaryExpr(e Unary[string]) string {
	return ap.parenthesize(e.Operator.Lexeme, e.Exp)
}

func (ap *AstPrinter) VisitGroupingExpr(e Grouping[string]) string {
	return ap.parenthesize("group", e.Exp)
}

func (ap *AstPrinter) VisitLiteralExpr(e Literal[string]) string {
	if e.Val == nil {
		return "null"
	}
	return fmt.Sprintf("%v", e.Val)
}

func (ap *AstPrinter) Print(e Expression[string]) string {
	return e.Accept(ap)
}

func (ap *AstPrinter) parenthesize(name string, exps ...Expression[string]) string {
	var sb strings.Builder
	sb.WriteRune('(')
	sb.WriteString(name)
	for _, exp := range exps {
		sb.WriteRune(' ')
		sb.WriteString(exp.Accept(ap))
	}
	sb.WriteRune(')')
	return sb.String()
}
