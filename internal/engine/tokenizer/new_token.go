package tokenizer

import "github.com/apastori/monkey-toto-interpreter-go/internal/domain/token"

func newToken(tokenType token.TokenType, ch byte) token.Token {
	var literal string = string(ch)
	var newTokenP *token.Token = new(token.Token)
	newTokenP.Type = tokenType
	newTokenP.Literal = literal
	var newToken token.Token = *newTokenP
	return newToken
}
