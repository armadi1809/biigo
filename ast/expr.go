package ast

import "github.com/armadi1809/biigo/token"

type Visitor[R any] interface {
	VisitBinaryExpr(Binary[R]) R
	VisitGroupingExpr(Grouping[R]) R
	VisitUnaryExpr(Unary[R]) R
	VisitLiteralExpr(Literal[R]) R
}

type Expression[R any] interface {
	Accept(Visitor[R]) R
}

type Binary[R any] struct {
	Left     Expression[R]
	Right    Expression[R]
	Operator token.Token
}

type Unary[R any] struct {
	Exp      Expression[R]
	Operator token.Token
}

type Grouping[R any] struct {
	Exp Expression[R]
}

type Literal[R any] struct {
	Val any
}

func (e Binary[R]) Accept(v Visitor[R]) R {
	return v.VisitBinaryExpr(e)
}

func (e Unary[R]) Accept(v Visitor[R]) R {
	return v.VisitUnaryExpr(e)
}

func (e Grouping[R]) Accept(v Visitor[R]) R {
	return v.VisitGroupingExpr(e)
}

func (e Literal[R]) Accept(v Visitor[R]) R {
	return v.VisitLiteralExpr(e)
}
