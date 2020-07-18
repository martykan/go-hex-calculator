package main

func operationPriority(token *Token) int {
	if token.value == "*" || token.value == "/" {
		return 1
	}
	return 0
}

func popNode(a *[]TreeNode) (x TreeNode) {
	x = (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return x
}

func addNode(operatorToken *Token, outputStack *[]TreeNode) {
	rightNode := popNode(outputStack)
	leftNode := popNode(outputStack)
	node := TreeNode{*operatorToken, &leftNode, &rightNode}
	*outputStack = append(*outputStack, node)
}

// ParseInput - parses the provided string input into a tree structure
func ParseInput(input string) *TreeNode {
	outputStack := []TreeNode{}
	operationStack := []*Token{}

	tokens := Tokenize(input)
	for _, token := range tokens {
		switch {
		case token.tokenType == NUMBER:
			node := TreeNode{*token, nil, nil}
			outputStack = append(outputStack, node)
		case token.tokenType == OPERATOR:
			for len(operationStack) > 0 {
				operation := operationStack[len(operationStack)-1]
				if operation.tokenType == OPERATOR && operationPriority(token) <= operationPriority(operation) {
					operationStack = operationStack[:len(operationStack)-1]
					addNode(operation, &outputStack)
				} else {
					break
				}
			}
			operationStack = append(operationStack, token)
		case token.tokenType == NONE && token.value == "(":
			operationStack = append(operationStack, token)
		case token.tokenType == NONE && token.value == ")":
			if len(operationStack) == 0 {
				panic("Unmatched parenthesis")
			}

			for len(operationStack) > 0 {
				operation := operationStack[len(operationStack)-1]
				if operation.value == "(" {
					break
				}
				operationStack = operationStack[:len(operationStack)-1]
				addNode(operation, &outputStack)
			}

			// Remove the left parenthesis from the stack
			operationStack = operationStack[:len(operationStack)-1]
		}
	}

	for len(operationStack) > 0 {
		operation := operationStack[len(operationStack)-1]
		operationStack = operationStack[:len(operationStack)-1]
		addNode(operation, &outputStack)
	}

	output := popNode(&outputStack)
	return &output
}
