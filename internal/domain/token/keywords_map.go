package token

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"var":      VAR,
	"const":    CONST,
}
