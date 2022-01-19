package main

import "fmt"

func main() {
	l1 := ListNode{Val: 1, Next: &ListNode{Val: 4, Next: nil}}
	l2 := ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 7, Next: nil}}}

	result := mergeTwoLists(&l1, &l2)

	if result != nil {
		for {
			fmt.Println(result.Val)
			if result.Next == nil {
				break
			}

			result = result.Next
		}
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

type ListNode struct {
	Val  int
	Next *ListNode
}
