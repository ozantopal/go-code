package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}

	fmt.Println(isSymmetric(root))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSameTree(mirrorTree(root.Left), root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}

		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	mirrorTree(root.Left)
	mirrorTree(root.Right)

	root.Left, root.Right = root.Right, root.Left

	return root
}
