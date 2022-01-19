package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

func List2Ints(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)

		if head.Next != nil {
			head = head.Next
		} else {
			head = nil
		}
	}

	return result
}

func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	result := &ListNode{}
	i := result
	for _, val := range nums {
		i.Next = &ListNode{Val: val}
		i = i.Next
	}

	return result.Next
}
