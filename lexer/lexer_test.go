package lexer

import (
	"interpreter-mod/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `+={}();,`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] Token type is not expected token type. expected %q, got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] Token literal is not expected literal. expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
