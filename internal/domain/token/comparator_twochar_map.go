package token

var ComparatorTwoCharMap = map[string]TokenType{
	"==": EQUAL,
	">=": GREATER_THAN_OR_EQUAL,
	"<=": LESS_THAN_OR_EQUAL,
	"!=": NOT_EQUAL,
	"<>": NOT_EQUAL_2,
}
