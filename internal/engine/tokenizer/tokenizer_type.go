package tokenizer

type Tokenizer struct {
	input           string
	currentPosition int
	nextPosition    int
	currentChar     byte
}
