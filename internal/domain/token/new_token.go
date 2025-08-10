package token

func newToken(tokenType TokenType, ch byte) Token {
	var literal string = string(ch)
	var newTokenP *Token = new(Token)
	newTokenP.Type = tokenType
	newTokenP.Literal = literal
	var newToken Token = *newTokenP
	return newToken
}
