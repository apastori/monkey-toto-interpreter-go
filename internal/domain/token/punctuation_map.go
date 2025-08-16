package token

var PunctuationMap = map[byte]TokenType{
	';': SEMICOLON,
	'(': LPAREN,
	')': RPAREN,
	',': COMMA,
	'{': LBRACE,
	'}': RBRACE,
	']': RBRACKET,
	'[': LBRACKET,
	':': COLON,
}
