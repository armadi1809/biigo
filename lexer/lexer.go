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
	return []rune(lexer.source)[lexer.current-1]
}

func (lexer *Lexer) addToken(tok token.TokenType, literal string) {
	lexeme := lexer.source[lexer.start:lexer.current]
	lexer.tokens = append(lexer.tokens, *token.NewToken(tok, lexeme, literal, lexer.line))
}
