package token

var PunctuationMap = map[byte]TokenType{
	';': SEMI_COLON,
	'(': L_PAREN,
	')': R_PAREN,
	',': COMMA,
	'{': L_BRACE,
	'}': R_BRACE,
	']': R_BRACKET,
	'[': L_BRACKET,
	':': COLON,
}
