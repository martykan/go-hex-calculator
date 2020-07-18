package main

import (
	"fmt"
	"os"
	"strings"
)

func printTree(depth int, node *TreeNode) {
	if node == nil {
		return
	}

	printTree(depth+1, node.leftNode)
	fmt.Println(strings.Repeat(" ", depth) + node.Token.value)
	printTree(depth+1, node.rightNode)
}

func main() {
	root := ParseInput(os.Args[1])
	printTree(0, root)
	result := EvaluateTree(root)
	fmt.Printf("\nResult: %s\n", result)
}
