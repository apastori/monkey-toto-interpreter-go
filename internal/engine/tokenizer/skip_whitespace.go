package tokenizer

func (tokenizer *Tokenizer) skipWhitespace() {
	for tokenizer.currentChar == ' ' || tokenizer.currentChar == '\t' || tokenizer.currentChar == '\n' || tokenizer.currentChar == '\r' {
		tokenizer.readChar()
	}
}
