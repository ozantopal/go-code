package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil},
			Right: nil}}
	fmt.Println(inorderTraversal(root))
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	stack := []*TreeNode{}
	visited := make(map[*TreeNode]bool)

	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		for node.Left != nil && !visited[node.Left] {
			node = node.Left
			stack = append(stack, node)
		}

		visited[node] = true
		stack = stack[0 : len(stack)-1]
		result = append(result, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
