package tokenizer

func (tokenizer *Tokenizer) readIdentifier() string {
	position := tokenizer.currentPosition
	for isLetter(tokenizer.currentChar) {
		tokenizer.readChar()
	}
	return tokenizer.input[position:tokenizer.currentPosition]
}
