package main

import (
	"strings"
)

// Functions to check token type
func isOperator(input rune) bool {
	return input == '+' || input == '-' || input == '*' || input == '/'
}

func isNumber(input rune) bool {
	return input >= 97 && input <= 102 || input >= 48 && input <= 57
}

func isLParenthesis(input rune) bool {
	return input == '('
}

func isRParenthesis(input rune) bool {
	return input == ')'
}

// Complete number buffer
func finishNumberBuffer(numberBuffer *[]rune, tokens *[]*Token) {
	if len(*numberBuffer) == 0 {
		return
	}
	addToken(NUMBER, string(*numberBuffer), tokens)
	*numberBuffer = []rune{}
}

// Helper function to add a new token
func addToken(tokenType TokenType, value string, tokens *[]*Token) {
	token := &Token{tokenType, value}
	*tokens = append(*tokens, token)
}

// Tokenize - generate an array of tokens from the input string
func Tokenize(input string) []*Token {
	tokens := []*Token{}
	// Remove whitespace
	input = strings.Replace(input, " ", "", -1)
	// To lowercase
	input = strings.ToLower(input)

	// Init number buffer
	numberBuffer := []rune{}

	for i, char := range input {
		switch {
		case isOperator(char):
			if rune(char) == '-' && (i == 0 || !isNumber(rune(input[i-1]))) {
				numberBuffer = append(numberBuffer, char)
			} else {
				finishNumberBuffer(&numberBuffer, &tokens)
				addToken(OPERATOR, string(char), &tokens)
			}
		case isNumber(char):
			numberBuffer = append(numberBuffer, char)
		case isLParenthesis(char):
			if len(numberBuffer) > 0 {
				finishNumberBuffer(&numberBuffer, &tokens)
				addToken(OPERATOR, "*", &tokens)
			}
			addToken(NONE, "(", &tokens)
		case isRParenthesis(char):
			finishNumberBuffer(&numberBuffer, &tokens)
			addToken(NONE, ")", &tokens)
		default:
			panic("Invalid input token")
		}
	}
	if len(numberBuffer) > 0 {
		finishNumberBuffer(&numberBuffer, &tokens)
	}
	return tokens
}
