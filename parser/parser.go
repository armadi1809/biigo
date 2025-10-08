package parser

import (
	"fmt"

	"github.com/armadi1809/biigo/ast"
	"github.com/armadi1809/biigo/langerror"
	"github.com/armadi1809/biigo/token"
)

type Parser struct {
	tokens  []token.Token
	current int
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) Parse() (ast.Expression, error) {
	return p.expression()
}

func (p *Parser) expression() (ast.Expression, error) {
	return p.equality()
}

func (p *Parser) equality() (ast.Expression, error) {
	e, err := p.comparison()
	if err != nil {
		return nil, err
	}
	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		e = ast.Binary{
			Left:     e,
			Operator: operator,
			Right:    right,
		}
	}
	return e, nil
}

func (p *Parser) comparison() (ast.Expression, error) {
	e, err := p.term()
	if err != nil {
		return nil, err
	}
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		e = ast.Binary{
			Left:     e,
			Operator: operator,
			Right:    right,
		}
	}
	return e, nil
}

func (p *Parser) term() (ast.Expression, error) {
	e, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.match(token.PLUS, token.MINUS) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		e = ast.Binary{
			Left:     e,
			Operator: operator,
			Right:    right,
		}
	}
	return e, nil
}

func (p *Parser) factor() (ast.Expression, error) {
	e, err := p.unary()
	if err != nil {
		return nil, err
	}
	for p.match(token.STAR, token.SLASH) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		e = ast.Binary{
			Left:     e,
			Operator: operator,
			Right:    right,
		}
	}
	return e, nil
}

func (p *Parser) unary() (ast.Expression, error) {
	if p.match(token.BANG, token.BANG_EQUAL) {
		operator := p.previous()
		e, err := p.unary()
		if err != nil {
			return nil, err
		}
		return ast.Unary{Operator: operator, Exp: e}, nil
	}

	return p.primary()
}

func (p *Parser) primary() (ast.Expression, error) {
	if p.match(token.FALSE) {
		return ast.Literal{Val: false}, nil
	}
	if p.match(token.TRUE) {
		return ast.Literal{Val: true}, nil
	}
	if p.match(token.NIL) {
		return ast.Literal{Val: nil}, nil
	}

	if p.match(token.NUMBER, token.STRING) {
		return ast.Literal{Val: p.previous().Literal}, nil
	}

	if p.match(token.LEFT_PAREN) {
		e, err := p.expression()
		if err != nil {
			return nil, err
		}
		err = p.consume(token.RIGHT_PAREN, "")
		if err != nil {
			return nil, err
		}
		return ast.Grouping{Exp: e}, nil
	}

	err := &langerror.LangError{
		Message: "parse error: expected an expression",
		Where:   "Parser",
		Line:    p.peek().Line,
	}
	return nil, err
}

func (p *Parser) consume(tok token.TokenType, errMessage string) error {
	if p.check(tok) {
		p.advance()
		return nil
	}
	err := &langerror.LangError{
		Message: fmt.Sprintf("parse error: %s", errMessage),
		Where:   "Parser",
		Line:    p.peek().Line,
	}
	return err
}

func (p *Parser) match(types ...token.TokenType) bool {
	for _, tokenType := range types {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(tokenType token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current += 1
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}
