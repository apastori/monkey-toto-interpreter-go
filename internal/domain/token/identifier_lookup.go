package token

func IdentifierLookup(identifier string) TokenType {
	var tokenKeywordIdent TokenType
	var exists bool
	// Checking punctuation
	if tokenKeywordIdent, exists = keywords[identifier]; exists {
		return tokenKeywordIdent
	}
	// Return nil when no token is found
	return IDENTIFIER
}
