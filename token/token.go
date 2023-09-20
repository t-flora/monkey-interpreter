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
	PLUS     = "+"
	ASSIGN   = "="
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"
	LT       = "<"
	GT       = ">"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
