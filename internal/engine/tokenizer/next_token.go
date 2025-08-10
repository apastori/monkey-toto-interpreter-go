package tokenizer

import "github.com/apastori/monkey-toto-interpreter-go/internal/domain/token"

func (tokenizer *Tokenizer) NextToken() token.Token {
	// Handle EOF case first
	if tokenizer.currentChar == 0 {
		tokenizer.readChar()
		return token.Token{
			Type:    token.EOF,
			Literal: "",
		}
	}
	var token token.Token
	token = tokenLookup(tokenizer.currentChar)
	tokenizer.readChar()
	return token
}
