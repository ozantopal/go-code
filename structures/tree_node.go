package structures

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		result = append(result, root.Val)

		result = append(result, PreorderTraversal(root.Left)...)
		result = append(result, PreorderTraversal(root.Right)...)
	}

	return result
}
