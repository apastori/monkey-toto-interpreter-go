package token

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"
	// Identifiers + literals
	IDENTIFIER TokenType = "IDENTIFIER"
	INT        TokenType = "INT"    // 1
	STRING     TokenType = "STRING" // "hello"
	// Operators
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"
	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	// Keywords
	VAR      TokenType = "VAR"
	CONST    TokenType = "CONST"
	FUNCTION TokenType = "FUNCTION"
)
