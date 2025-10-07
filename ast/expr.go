package ast

import (
	"fmt"
	"strings"

	"github.com/armadi1809/biigo/token"
)

type Expression interface {
	String() string
}

type Binary struct {
	Left     Expression
	Right    Expression
	Operator token.Token
}

type Unary struct {
	Exp      Expression
	Operator token.Token
}

type Grouping struct {
	Exp Expression
}

type Literal struct {
	Val any
}

func (e Binary) String() string {
	return parenthesize(e.Operator.Lexeme, e.Left, e.Right)
}

func (e Unary) String() string {
	return parenthesize(e.Operator.Lexeme, e.Exp)
}

func (e Grouping) String() string {
	return parenthesize("group", e.Exp)
}

func (e Literal) String() string {
	if e.Val == nil {
		return "null"
	}
	return fmt.Sprintf("%v", e.Val)
}

func parenthesize(name string, exps ...Expression) string {
	var sb strings.Builder
	sb.WriteRune('(')
	sb.WriteString(name)
	for _, exp := range exps {
		sb.WriteRune(' ')
		sb.WriteString(exp.String())
	}
	sb.WriteRune(')')
	return sb.String()
}
