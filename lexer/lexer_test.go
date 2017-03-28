package lexer

import (
	"testing"

	"github.com/bchip/trippi-CS451/token"
)

func TestNextToken(t *testing.T) {
	input := `
			2.25
			const
			input
			let five = 5;
			let ten = 10;
			if (5 < 10) {
			} else {
				let t = true;
			}
			`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.FLOAT, "2.25"},
		{token.CONST, "const"},
		{token.INPUT, "input"},
		{token.LET, "let"}, {token.IDENT, "five"},
		{token.ASSIGN, "="}, {token.INT, "5"}, {token.SEMICOLON, ";"}, {token.LET, "let"},
		{token.IDENT, "ten"}, {token.ASSIGN, "="}, {token.INT, "10"}, {token.SEMICOLON, ";"},
		{token.IF, "if"}, {token.LPAREN, "("},
		{token.INT, "5"}, {token.LT, "<"}, {token.INT, "10"}, {token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"}, {token.ELSE, "else"}, {token.LBRACE, "{"},
		{token.LET, "let"}, {token.IDENT, "t"}, {token.ASSIGN, "="}, {token.TRUE, "true"}, {token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
