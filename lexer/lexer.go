package lexer

import "github.com/armadi1809/biigo/token"

type Lexer struct {
	source string
}

func NewLexer(source string) *Lexer {
	return &Lexer{
		source: source,
	}
}

func (lexer *Lexer) ScanTokens() []token.Token {
	return []token.Token{}
}
