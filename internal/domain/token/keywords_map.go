package token

var keywords = map[string]TokenType{
	"func": FUNCTION,
	"var":      VAR,
	"const":    CONST,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"fi":	   FI,
	"else":     ELSE,
	"elif":     ELIF,
	"return":   RETURN,
}
