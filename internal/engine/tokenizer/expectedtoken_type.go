package tokenizer

import "github.com/apastori/monkey-toto-interpreter-go/internal/domain/token"

type expectedToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}
