package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Lexeme  string
	Line    int
}

const (
	EOF    = "EOF"
	STRING = "STRING"
	NUMBER = "NUMBER"

	EQUAL = "="
	PLUS  = "+"
	MINUS = "-"
	SLASH = "/"
	STAR  = "*"

	BANG          = "!"
	BANG_EQUAL    = "!="
	EQUAL_EQUAL   = "=="
	GREATER       = ">"
	GREATER_EQUAL = ">="
	LESS          = "<"
	LESS_EQUAL    = "<="

	COMMA     = ","
	SEMICOLON = ";"
	DOT       = "."

	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"

	// keywords
	AND      = "AND"
	CLASS    = "CLASS"
	ELSE     = "ELSE"
	FALSE    = "FALSE"
	FUN      = "FUN"
	FOR      = "FOR"
	IF       = "IF"
	NIL      = "NIL"
	OR       = "OR"
	PRINT    = "PRINT"
	RETURN   = "RETURN"
	SUPER    = "SUPER"
	THIS     = "THIS"
	TRUE     = "TRUE"
	VAR      = "VAR"
	WHILE    = "WHILE"
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

func NewToken(t TokenType, lexeme string, literal string, line int) *Token {
	return &Token{
		Type:    t,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}
