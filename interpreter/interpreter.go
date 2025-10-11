package interpreter

import (
	"fmt"

	"github.com/armadi1809/biigo/ast"
	"github.com/armadi1809/biigo/langerror"
	"github.com/armadi1809/biigo/token"
)

func Interpret(e ast.Expression) (any, error) {
	switch e := (e).(type) {
	case ast.Literal:
		return e.Val, nil
	case ast.Grouping:
		return Interpret(e.Exp)
	case ast.Unary:
		return evaluateUnary(e)
	case ast.Binary:
		return evaluateBinary(e)
	}

	err := &langerror.LangError{
		Message: fmt.Sprintf("invalid expression %s", e),
	}
	return nil, err
}

func evaluateBinary(e ast.Binary) (any, error) {
	lval, err := Interpret(e.Left)
	if err != nil {
		return nil, err
	}

	rval, err := Interpret(e.Right)
	if err != nil {
		return nil, err
	}

	langErr := &langerror.LangError{}

	switch e.Operator.Type {
	case token.PLUS:
		lNum, leftIsFloat := lval.(float64)
		rNum, rightIsFloat := rval.(float64)
		if leftIsFloat && rightIsFloat {
			return lNum + rNum, nil
		}
		lstr, leftIsString := lval.(string)
		rstr, rightIsString := rval.(string)
		if leftIsString && rightIsString {
			return lstr + rstr, nil
		}
		langErr.Line = e.Operator.Line
		langErr.Message = fmt.Sprintf("Invalid operands %v, %v for operator %s", lval, rval, e.Operator.Type)
		return nil, langErr
	case token.MINUS:
		lNum, leftIsFloat := lval.(float64)
		rNum, rightIsFloat := rval.(float64)
		if leftIsFloat && rightIsFloat {
			return lNum - rNum, nil
		}
		langErr.Line = e.Operator.Line
		langErr.Message = fmt.Sprintf("Invalid operands %v, %v for operator %s", lval, rval, e.Operator.Type)
		return nil, langErr
	case token.STAR:
		lNum, leftIsFloat := lval.(float64)
		rNum, rightIsFloat := rval.(float64)
		if leftIsFloat && rightIsFloat {
			return lNum * rNum, nil
		}
		langErr.Line = e.Operator.Line
		langErr.Message = fmt.Sprintf("Invalid operands %v, %v for operator %s", lval, rval, e.Operator.Type)
		return nil, langErr
	case token.SLASH:
		lNum, leftIsFloat := lval.(float64)
		rNum, rightIsFloat := rval.(float64)
		if leftIsFloat && rightIsFloat {
			return lNum / rNum, nil
		}
		langErr.Line = e.Operator.Line
		langErr.Message = fmt.Sprintf("Invalid operands %v, %v for operator %s", lval, rval, e.Operator.Type)
		return nil, langErr
	case token.GREATER:
		lNum, leftIsFloat := lval.(float64)
		rNum, rightIsFloat := rval.(float64)
		if leftIsFloat && rightIsFloat {
			return lNum > rNum, nil
		}
		langErr.Line = e.Operator.Line
		langErr.Message = fmt.Sprintf("Invalid operands %v, %v for operator %s", lval, rval, e.Operator.Type)
		return nil, langErr
	case token.GREATER_EQUAL:
		lNum, leftIsFloat := lval.(float64)
		rNum, rightIsFloat := rval.(float64)
		if leftIsFloat && rightIsFloat {
			return lNum >= rNum, nil
		}
		langErr.Line = e.Operator.Line
		langErr.Message = fmt.Sprintf("Invalid operands %v, %v for operator %s", lval, rval, e.Operator.Type)
		return nil, langErr
	case token.LESS:
		lNum, leftIsFloat := lval.(float64)
		rNum, rightIsFloat := rval.(float64)
		if leftIsFloat && rightIsFloat {
			return lNum < rNum, nil
		}
		langErr.Line = e.Operator.Line
		langErr.Message = fmt.Sprintf("Invalid operands %v, %v for operator %s", lval, rval, e.Operator.Type)
		return nil, langErr
	case token.LESS_EQUAL:
		lNum, leftIsFloat := lval.(float64)
		rNum, rightIsFloat := rval.(float64)
		if leftIsFloat && rightIsFloat {
			return lNum <= rNum, nil
		}
		langErr.Line = e.Operator.Line
		langErr.Message = fmt.Sprintf("Invalid operands %v, %v for operator %s", lval, rval, e.Operator.Type)
		return nil, langErr
	case token.BANG_EQUAL:
		return lval != rval, nil
	case token.EQUAL_EQUAL:
		return lval == rval, nil

	}

	return nil, nil
}

func evaluateUnary(e ast.Unary) (any, error) {
	val, err := Interpret(e.Exp)
	if err != nil {
		return nil, err
	}
	switch e.Operator.Type {
	case token.BANG:
		return !isTruthy(val), nil
	case token.MINUS:
		valNum, ok := val.(float64)
		if !ok {
			err = &langerror.LangError{
				Message: fmt.Sprintf("invalid operand %v for operator %s", val, e.Operator.Type),
				Line:    e.Operator.Line,
			}
			return nil, err
		}
		return -valNum, nil
	}
	return nil, nil
}

func isTruthy(val any) bool {
	if val == nil {
		return false
	}
	if v, ok := val.(bool); ok {
		return v
	}
	return true
}
