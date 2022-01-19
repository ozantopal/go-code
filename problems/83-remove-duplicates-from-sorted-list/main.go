package main

import "fmt"

func main() {
	x := ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}}
	result := deleteDuplicates(&x)
	for {
		fmt.Println(result.Val)
		if result.Next == nil {
			break
		}
		result = result.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	result := head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		}

		head = head.Next
	}

	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}
