package lexer

import (
	"log"

	"github.com/armadi1809/biigo/langerror"
	"github.com/armadi1809/biigo/token"
)

type Lexer struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
}

func NewLexer(source string) *Lexer {
	return &Lexer{
		source:  source,
		tokens:  []token.Token{},
		start:   0,
		current: 0,
		line:    0,
	}
}

func (lexer *Lexer) ScanTokens() ([]token.Token, error) {

	for !lexer.isAtEnd() {
		lexer.start = lexer.current
		err := lexer.scanToken()
		if err != nil {
			log.Println(err.Error())
		}
	}
	lexer.tokens = append(lexer.tokens, *token.NewToken(token.EOF, "", "", lexer.line))
	return lexer.tokens, nil
}

func (lexer *Lexer) isAtEnd() bool {
	return lexer.current >= len(lexer.source)
}

func (lexer *Lexer) scanToken() error {
	c := lexer.advance()
	var err error = nil
	switch c {
	case '(':
		lexer.addToken(token.LEFT_PAREN, "")
	case ')':
		lexer.addToken(token.RIGHT_PAREN, "")
	case '{':
		lexer.addToken(token.LEFT_BRACE, "")
	case '}':
		lexer.addToken(token.RIGHT_BRACE, "")
	case ',':
		lexer.addToken(token.COMMA, "")
	case '.':
		lexer.addToken(token.DOT, "")
	case '-':
		lexer.addToken(token.MINUS, "")
	case '+':
		lexer.addToken(token.PLUS, "")
	case ';':
		lexer.addToken(token.SEMICOLON, "")
	case '*':
		lexer.addToken(token.STAR, "")
	case '!':
		if lexer.match('=') {
			lexer.addToken(token.BANG_EQUAL, "")
		} else {
			lexer.addToken(token.BANG, "")
		}
	case '<':
		if lexer.match('=') {
			lexer.addToken(token.LESS_EQUAL, "")
		} else {
			lexer.addToken(token.LESS, "")
		}
	case '>':
		if lexer.match('=') {
			lexer.addToken(token.GREATER_EQUAL, "")
		} else {
			lexer.addToken(token.GREATER, "")
		}
	case '/':
		if lexer.match('/') {
			for lexer.peek() != '\n' && !lexer.isAtEnd() {
				lexer.advance()
			}
		} else {
			lexer.addToken(token.SLASH, "")
		}
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		lexer.line += 1

	default:
		err = &langerror.LangError{
			Line:    lexer.line,
			Message: "Unexpected character",
		}
	}
	return err
}

func (lexer *Lexer) advance() int32 {
	lexer.current += 1
	return lexer.getCharAtPos(lexer.current - 1)
}

func (lexer *Lexer) addToken(tok token.TokenType, literal string) {
	lexeme := lexer.source[lexer.start:lexer.current]
	lexer.tokens = append(lexer.tokens, *token.NewToken(tok, lexeme, literal, lexer.line))
}

func (lexer *Lexer) match(expected int32) bool {
	if lexer.isAtEnd() {
		return false
	}
	if lexer.getCharAtPos(lexer.current) != expected {
		return false
	}
	lexer.current += 1
	return true
}

func (lexer *Lexer) peek() int32 {
	if lexer.isAtEnd() {
		return '\000'
	}
	return lexer.getCharAtPos(lexer.current)
}

func (lexer *Lexer) getCharAtPos(pos int) int32 {
	return []rune(lexer.source)[pos]
}
