package tokenizer

func (tokenizer *Tokenizer) peekChar() byte {
	if tokenizer.nextPosition >= len(tokenizer.input) {
		return 0
	} else {
		return tokenizer.input[tokenizer.nextPosition]
	}
}
