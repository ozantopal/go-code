package main

import (
	"fmt"

	"github.com/ozantopal/go-code/structures"
)

type TreeNode = structures.TreeNode

func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	fmt.Println(postorderTraversal(root))
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		result = append(result, postorderTraversal(root.Left)...)
		result = append(result, postorderTraversal(root.Right)...)

		result = append(result, root.Val)
	}

	return result
}
