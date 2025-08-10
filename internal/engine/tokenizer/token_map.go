package tokenizer

import "github.com/apastori/monkey-toto-interpreter-go/internal/domain/token"

var tokenMap = map[byte]token.TokenType{
	'=': token.ASSIGN,
	';': token.SEMICOLON,
	'(': token.LPAREN,
	')': token.RPAREN,
	',': token.COMMA,
	'+': token.PLUS,
	'{': token.LBRACE,
	'}': token.RBRACE,
}
