package tokenizer

import "github.com/apastori/monkey-toto-interpreter-go/internal/domain/token"

func tokenLookup(ch byte) token.Token {
	var tokenTypeMap token.TokenType
	var exists bool
	if tokenTypeMap, exists = tokenMap[ch]; exists {
		return newToken(tokenTypeMap, ch)
	}
	return token.Token{
		Type:    token.ILLEGAL,
		Literal: string(ch),
	} // Default ILLEGAL token
}
