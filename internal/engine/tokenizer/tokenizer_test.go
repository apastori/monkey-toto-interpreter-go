package tokenizer

import (
	"github.com/apastori/monkey-toto-interpreter-go/internal/domain/token"
	"testing"
)

func TestTokenizer(testingContext *testing.T) {
	var input string = `var five = 5;
				var ten = 10;
				var add = function(x, y) {
				x + y;
				};
				var result = add(five, ten);
			`
	var testsTokens []expectedToken = []expectedToken{
		{token.VAR, "var"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "function"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
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
