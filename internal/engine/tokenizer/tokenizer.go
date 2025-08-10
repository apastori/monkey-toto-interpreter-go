package tokenizer

func NewTokenizer(input string) *Tokenizer {
	var newTokenizer *Tokenizer = &Tokenizer{input: input}
	if len(input) > 0 {
		newTokenizer.currentChar = input[0]
		newTokenizer.currentPosition = 0
		newTokenizer.nextPosition = 1
	}
	return newTokenizer
}
