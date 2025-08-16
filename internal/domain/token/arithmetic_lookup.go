package token

func ArithmeticLookup(ch byte) *Token {
	var tokenTypeMap TokenType
	var exists bool
	var tokenPtr *Token
	// Checking punctuation
	if tokenTypeMap, exists = ArithmeticMap[ch]; exists {
		tokenFound := newToken(tokenTypeMap, ch)
		tokenPtr = &tokenFound
		return tokenPtr
	}
	// Return nil when no token is found
	return nil
}
