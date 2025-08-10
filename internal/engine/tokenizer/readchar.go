package tokenizer

func (tokenizer *Tokenizer) readChar() {
	if tokenizer.nextPosition >= len(tokenizer.input) {
		tokenizer.currentChar = 0
		tokenizer.currentPosition = tokenizer.nextPosition
		tokenizer.nextPosition += 1
		return
	}
	tokenizer.currentChar = tokenizer.input[tokenizer.nextPosition]
	tokenizer.currentPosition = tokenizer.nextPosition
	tokenizer.nextPosition += 1
}
