package main

// TokenType - represents different token types
type TokenType int

const (
	// NONE - does nothing, used for grouping
	NONE TokenType = iota
	// NUMBER - represents a numerical value
	NUMBER
	// OPERATOR - represents an operation
	OPERATOR
)

// Token -
type Token struct {
	tokenType TokenType
	value     string
}
