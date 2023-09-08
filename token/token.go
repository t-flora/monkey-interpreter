package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// define token types
const (
	// special types
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// operators
	PLUS   = "+"
	ASSIGN = "="

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
