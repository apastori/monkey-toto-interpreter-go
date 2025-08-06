package tokenizer

import (
	"testing"
	"github.com/apastori/monkey-toto-interpreter-go/token"
)

tests_tokens := []struct {
expectedType token.TokenType
expectedLiteral string
}{
{token.ASSIGN, "="},
{token.PLUS, "+"},
{token.LPAREN, "("},
{token.RPAREN, ")"},
{token.LBRACE, "{"},
{token.RBRACE, "}"},
{token.COMMA, ","},
{token.SEMICOLON, ";"},
{token.EOF, ""},
}

new_tokenizer := New(input)
for index, test_token := range tests {
token := new_tokenizer.NextToken()
if token.Type != test_token.expectedType {
t.Fatalf("tests[%d] - token_type wrong. expected=%q, got=%q",
i, test_token.expectedType, token.Type)
}
if token.Literal != test_token.expectedLiteral {
test_token.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
i, test_token.expectedLiteral, token.Literal)
}
}



