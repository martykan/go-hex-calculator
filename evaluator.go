package main

import (
	"strings"
)

// Converts char hex representation to int value
// Fails quitely - returns 0
func hexCharToDec(input rune) int {
	if input >= 97 && input <= 102 {
		return int(input) - 87
	} else if input >= 48 && input <= 57 {
		return int(input) - 48
	} else {
		return 0
	}
}

// Converts int value to a hex char + decimal carry
func decToHexChar(input int) (carry int, output rune) {
	carry = input / 16
	input = (input % 16)
	if input < 0 {
		input = 16 + input
		carry = -1
	}
	if input >= 10 {
		output = rune(input + 87)
	} else {
		output = rune(input + 48)
	}
	return
}

func multiplication(a string, b string) (negative bool, result string) {
	negative = (strings.Index(a, "-") == 0) != (strings.Index(b, "-") == 0)
	a = strings.Replace(a, "-", "", -1)
	b = strings.Replace(b, "-", "", -1)

	aLen := len(a)
	a = strings.Repeat("0", len(b)) + a
	b = strings.Repeat("0", aLen) + b

	result = ""

	var carry int
	lastOutput := "0"
	for x := len(a) - 1; x >= 0; x-- {
		output := ""
		for y := len(a) - 1; y >= 0; y-- {
			aDec := hexCharToDec(rune(a[x]))
			bDec := hexCharToDec(rune(b[y]))
			resultDec := bDec*aDec + carry
			carryV, resultHex := decToHexChar(resultDec)
			carry = carryV
			if y >= 0 || resultHex != '0' {
				output = string(resultHex) + output
			}
		}
		if carry != 0 {
			_, carryHex := decToHexChar(carry)
			output = string(carryHex) + output
		}
		if output == strings.Repeat("0", len(output)) {
			break
		}
		_, lastOutput = additionAndSubstraction(lastOutput, output+strings.Repeat("0", len(a)-x-1), "+")
	}
	result = lastOutput
	return
}

func divison(a string, b string) (negative bool, result string) {
	negative = (strings.Index(a, "-") == 0) != (strings.Index(b, "-") == 0)
	a = strings.Replace(a, "-", "", -1)
	b = strings.Replace(b, "-", "", -1)

	aLen := len(a)
	a = strings.Repeat("0", len(b)) + a
	b = strings.Repeat("0", aLen) + b

	for i := 0; i < len(a); i++ {
		aDec := hexCharToDec(rune(a[i]))
		bDec := hexCharToDec(rune(b[i]))
		if aDec < bDec {
			break
		} else if aDec > bDec {
			panic("No fractions allowed, only integers")
		}
	}

	if a == strings.Repeat("0", len(a)) {
		panic("Division by zero")
	}

	result = ""

	var resultDec = 0
	lastOutput := b
	for {
		negative, lastOutput = additionAndSubstraction(a, lastOutput, "-")
		resultDec++
		if negative || lastOutput == strings.Repeat("0", len(lastOutput)) {
			break
		}
	}
	for resultDec > 0 {
		remainder := resultDec % 16
		resultDec = resultDec / 16
		_, hex := decToHexChar(remainder)
		result = string(hex) + result
	}
	return
}

func additionAndSubstraction(a string, b string, operation string) (negative bool, output string) {
	aNegative := strings.Index(a, "-") == 0
	bNegative := strings.Index(b, "-") == 0
	a = strings.Replace(a, "-", "", -1)
	b = strings.Replace(b, "-", "", -1)

	if len(a) > len(b) {
		b = strings.Repeat("0", (len(a)-len(b))) + b
	} else {
		a = strings.Repeat("0", (len(b)-len(a))) + a
	}

	if operation == "+" && aNegative && bNegative {
		operation = "+"
		negative = true
	} else if operation == "+" && aNegative && !bNegative {
		operation = "-"
	} else if operation == "+" && !aNegative && bNegative {
		operation = "-"
		negative = true
	}

	invertNeg := false
	if operation == "-" && ((aNegative && bNegative) || (!aNegative && bNegative)) {
		invertNeg = true
	} else if operation == "-" && aNegative && !bNegative {
		operation = "+"
	}

	if operation == "-" {
		for i := 0; i < len(a); i++ {
			aDec := hexCharToDec(rune(a[i]))
			bDec := hexCharToDec(rune(b[i]))
			if aDec < bDec {
				break
			} else if aDec > bDec {
				tmp := b
				b = a
				a = tmp
				negative = true
				break
			}
		}
	}

	var carry int
	for i := len(a) - 1; i >= 0; i-- {
		aDec := hexCharToDec(rune(a[i]))
		bDec := hexCharToDec(rune(b[i]))
		resultDec := aDec
		if operation == "+" {
			resultDec = bDec + aDec + carry
		} else if operation == "-" {
			resultDec = bDec - aDec + carry
		}
		carryV, resultHex := decToHexChar(resultDec)
		carry = carryV
		if i >= 0 || resultHex != '0' {
			output = string(resultHex) + output
		}
	}
	if carry != 0 {
		_, carryHex := decToHexChar(carry)
		output = string(carryHex) + output
	}

	if invertNeg {
		negative = !negative
	}
	return
}

// EvaluateTree - Evaluates the provided tree
func EvaluateTree(node *TreeNode) (output string) {
	switch {
	case node.Token.tokenType == NUMBER:
		return node.Token.value
	default:
		a := EvaluateTree(node.rightNode)
		b := EvaluateTree(node.leftNode)
		negative := false
		if node.Token.value == "*" {
			negative, output = multiplication(a, b)
		} else if node.Token.value == "/" {
			negative, output = divison(a, b)
		} else {
			negative, output = additionAndSubstraction(a, b, node.Token.value)
		}
		for {
			pos := strings.Index(output, "0")
			if pos == 0 && pos < len(output)-1 {
				output = output[1:]
			} else {
				break
			}
		}
		if negative {
			output = "-" + output
		}
	}
	return
}
