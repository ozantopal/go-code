package main

import (
	"fmt"

	"github.com/ozantopal/go-code/structures"
)

type TreeNode = structures.TreeNode

func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	fmt.Println(preorderTraversal(root))
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		result = append(result, root.Val)

		result = append(result, preorderTraversal(root.Left)...)
		result = append(result, preorderTraversal(root.Right)...)
	}

	return result
}
