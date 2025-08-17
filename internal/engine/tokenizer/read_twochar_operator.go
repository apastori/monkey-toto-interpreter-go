package tokenizer

import "github.com/apastori/monkey-toto-interpreter-go/internal/domain/token"
// readTwoCharOperator checks if the current character and the next character form a two-character operator.

func (tokenizer *Tokenizer) readTwoCharOperator() *token.Token {
	var ch byte = tokenizer.currentChar
	if !isTwoCharCandidate(ch) {
		return nil	
	}
	var nextChar byte = tokenizer.peekChar()
	var literal string = string(ch) + string(nextChar)
	if tokType, ok := token.ComparatorTwoCharMap[literal]; ok {
		tokenizer.readChar() // consume current char
		tokenizer.readChar() // consume second char
		return &token.Token{Type: tokType, Literal: literal}
	}
	if tokType, ok := token.LogicalTwoCharMap[literal]; ok {
		tokenizer.readChar() // consume current char
		tokenizer.readChar() // consume second char
		return &token.Token{Type: tokType, Literal: literal}
	}
	return nil
}