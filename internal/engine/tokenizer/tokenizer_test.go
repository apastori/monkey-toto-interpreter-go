package tokenizer

import (
	"github.com/apastori/monkey-toto-interpreter-go/internal/domain/token"
	"testing"
)

func TestTokenizer(testingContext *testing.T) {
	var input string = `=+(){},;`
	var testsTokens []expectedToken = []expectedToken{
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
	var newTokenizer *Tokenizer = NewTokenizer(input)
	var index int
	var testToken expectedToken
	for index, testToken = range testsTokens {
		var nextToken token.Token = newTokenizer.NextToken()
		if nextToken.Type != testToken.expectedType {
			testingContext.Fatalf("tests[%d] - token_type wrong. expected=%q, got=%q",
				index, testToken.expectedType, nextToken.Type)
		}
		if nextToken.Literal != testToken.expectedLiteral {
			testingContext.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				index, testToken.expectedLiteral, nextToken.Literal)
		}
	}
}
