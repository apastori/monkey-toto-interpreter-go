package tokenizer

import (
	"testing"

	"github.com/apastori/monkey-toto-interpreter-go/internal/domain/token"
)

func TestTokenizer(testingContext *testing.T) {
	var input string = `var five = 5;
		var ten = 10;
		var add = func(x, y) {
			x + y;
		};
		var result = add(five, ten);
		!-/*5;
		5 < 10 > 5;
		if (5 < 10) {
			return true;
		} else {
			return false;
		}
		10 == 10;
		10 != 9;
		8 <> 9;
		`
	var testsTokens []expectedToken = []expectedToken{
		{token.VAR, "var"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMI_COLON, ";"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMI_COLON, ";"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "func"},
		{token.L_PAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.R_PAREN, ")"},
		{token.L_BRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMI_COLON, ";"},
		{token.R_BRACE, "}"},
		{token.SEMI_COLON, ";"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.L_PAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.R_PAREN, ")"},
		{token.SEMI_COLON, ";"},
		{token.NOT, "!"},
		{token.MINUS, "-"},
		{token.DIVIDE, "/"},
		{token.MULTIPLY, "*"},
		{token.INT, "5"},
		{token.SEMI_COLON, ";"},
		{token.INT, "5"},
		{token.LESS_THAN, "<"},
		{token.INT, "10"},
		{token.GREATER_THAN, ">"},
		{token.INT, "5"},
		{token.SEMI_COLON, ";"},
		{token.IF, "if"},
		{token.L_PAREN, "("},
		{token.INT, "5"},
		{token.LESS_THAN, "<"},
		{token.INT, "10"},
		{token.R_PAREN, ")"},
		{token.L_BRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMI_COLON, ";"},
		{token.R_BRACE, "}"},
		{token.ELSE, "else"},
		{token.L_BRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMI_COLON, ";"},
		{token.R_BRACE, "}"},
		{token.INT, "10"},
		{token.EQUAL, "=="},
		{token.INT, "10"},
		{token.SEMI_COLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQUAL, "!="},
		{token.INT, "9"},
		{token.SEMI_COLON, ";"},
		{token.INT, "8"},
		{token.NOT_EQUAL, "<>"},
		{token.INT, "9"},
		{token.SEMI_COLON, ";"},
		// EOF token
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
