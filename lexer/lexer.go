package lexer

import (
	"log"
	"strconv"

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

var keywords = map[string]token.TokenType{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"if":     token.IF,
	"false":  token.FALSE,
	"true":   token.TRUE,
	"var":    token.VAR,
	"fun":    token.FUN,
	"while":  token.WHILE,
	"super":  token.SUPER,
	"print":  token.PRINT,
	"nil":    token.NIL,
	"or":     token.OR,
	"return": token.RETURN,
	"for":    token.FOR,
	"this":   token.THIS,
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
	case '"':
		lit, err := lexer.stringLiteral()
		if err != nil {
			break
		}
		lexer.addToken(token.STRING, lit)

	default:
		if isDigit(c) {
			num, err := lexer.numberLiteral()
			if err != nil {
				break
			}
			lexer.addToken(token.NUMBER, num)
		} else if isAlpha(c) {
			iden := lexer.identifier()
			t, ok := keywords[iden]
			if !ok {
				t = token.IDENTIFIER
			}
			lexer.addToken(t, "")
		} else {
			err = &langerror.LangError{
				Line:    lexer.line,
				Message: "Unexpected character",
			}
		}
	}
	return err
}

func (lexer *Lexer) advance() int32 {
	lexer.current += 1
	return lexer.getCharAtPos(lexer.current - 1)
}

func (lexer *Lexer) addToken(tok token.TokenType, literal any) {
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

func (lexer *Lexer) peekNext() int32 {
	if lexer.current+1 >= len(lexer.source) {
		return '\000'
	}
	return lexer.getCharAtPos(lexer.current + 1)
}

func (lexer *Lexer) getCharAtPos(pos int) int32 {
	return []rune(lexer.source)[pos]
}

func (lexer *Lexer) stringLiteral() (string, error) {
	for lexer.peek() != '"' && !lexer.isAtEnd() {
		if lexer.peek() == '\n' {
			lexer.line += 1
		}
		lexer.advance()
	}
	if lexer.isAtEnd() {
		return "", &langerror.LangError{
			Line:    lexer.line,
			Message: "Unterminated string",
		}
	}
	lexer.advance()
	return lexer.source[lexer.start+1 : lexer.current-1], nil
}

func (lexer *Lexer) numberLiteral() (float64, error) {
	for isDigit(lexer.peek()) {
		lexer.advance()
	}
	if lexer.peek() == '.' && isDigit(lexer.peekNext()) {
		lexer.advance()
		for isDigit(lexer.peek()) {
			lexer.advance()
		}
	}
	num, err := strconv.ParseFloat(lexer.source[lexer.start:lexer.current], 64)
	if err != nil {
		return -1, &langerror.LangError{
			Line:    lexer.line,
			Message: "Invalid number",
		}
	}
	return num, nil
}

func (lexer *Lexer) identifier() string {
	for isAlphaNum(lexer.peek()) {
		lexer.advance()
	}

	return lexer.source[lexer.start:lexer.current]

}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isAlphaNum(c rune) bool {
	return isAlpha(c) || isDigit(c)
}
