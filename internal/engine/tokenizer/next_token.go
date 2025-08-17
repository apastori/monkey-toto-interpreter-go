package tokenizer

import (
	"github.com/apastori/monkey-toto-interpreter-go/internal/domain/token"
	"github.com/apastori/monkey-toto-interpreter-go/internal/utils"
)

func (tokenizer *Tokenizer) NextToken() token.Token {
	var newToken token.Token
	// Skip whitespace
	tokenizer.skipWhitespace()
	// Handle EOF case first
	if tokenizer.currentChar == 0 {
		// advances to next character
		tokenizer.readChar()
		return token.Token{
			Type:    token.EOF,
			Literal: "",
		}
	}
	// Check for two-character tokens first (Comparators, Logical Operators)
	if isTwoCharCandidate(tokenizer.currentChar) {
		// Check for two-character tokens with TwoCharLookup
		if tok := tokenizer.readTwoCharOperator(); tok != nil {
			return *tok
		}
	}
	// Check for one-char logical operators
	if token.IsLogicalOneChar(tokenizer.currentChar) {
		// Check logical operators with LogicalOneCharLookup
		tokenizer.readChar()
		return token.Token{
			Type:    token.NOT,
			Literal: "!",
		}
	}
	// check punctuation or delimiters with PunctuationLookup
	tokenPunctuationPtr := token.PunctuationLookup(tokenizer.currentChar)
	if tokenPunctuationPtr != nil {
		// advances to next character
		tokenizer.readChar()
		// dereference pointer, return just the value
		return *tokenPunctuationPtr
	}
	// check assignment "=" with AssignmentLookup
	tokenAssignmentPtr := token.AssignmentLookup(tokenizer.currentChar)
	if tokenAssignmentPtr != nil {
		// advances to next character
		tokenizer.readChar()
		// dereference pointer, return just the value
		return *tokenAssignmentPtr
	}
	// check arithmetic operations + - * / % with ArithmeticLookup
	tokenArithmeticPtr := token.ArithmeticLookup(tokenizer.currentChar)
	if tokenArithmeticPtr != nil {
		// advances to next character
		tokenizer.readChar()
		// dereference pointer, return just the value
		return *tokenArithmeticPtr
	}
	// check one-char comparisons > < with ComparisonOneCharLookup
	tokenComparisonOneCharPtr := token.ComparisonOneCharLookup(tokenizer.currentChar)
	if tokenComparisonOneCharPtr != nil {
		// advances to next character
		tokenizer.readChar()
		// dereference pointer, return just the value
		return *tokenComparisonOneCharPtr
	}
	// read keywords or identifier
	if utils.IsAlphabeticOrUnderscore(tokenizer.currentChar) {
		newToken.Literal = tokenizer.readIdentifier()
		newToken.Type = token.IdentifierLookup(newToken.Literal)
		return newToken
	}
	// read Numbers
	if utils.IsDigit(tokenizer.currentChar) {
		newToken.Literal = tokenizer.readNumber()
		newToken.Type = token.INT
		return newToken
	}
	// If we reach here, it means we have an illegal character
	newToken.Literal = string(tokenizer.currentChar)
	newToken.Type = token.ILLEGAL
	// advances to next character
	tokenizer.readChar()
	// Default case: return ILLEGAL token
	return newToken
}
