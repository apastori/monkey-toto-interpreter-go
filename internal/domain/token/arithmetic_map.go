package token

var ArithmeticMap = map[byte]TokenType{
	'+': PLUS,
	'-': MINUS,
	'*': MULTIPLY,
	'/': DIVIDE,
	'%': MODULO,
}
