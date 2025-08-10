package tokenizer

import "github.com/apastori/monkey-toto-interpreter-go/internal/utils"

func (tokenizer *Tokenizer) readNumber() string {
	position := tokenizer.currentPosition
	for utils.IsDigit(tokenizer.currentChar) {
		tokenizer.readChar()
	}
	return tokenizer.input[position:tokenizer.currentPosition]
}
