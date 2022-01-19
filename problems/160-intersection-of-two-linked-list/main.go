package main

import (
	"fmt"

	"github.com/ozantopal/go-code/structures"
)

type ListNode = structures.ListNode

func main() {
	c := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	a := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: c}}
	b := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: c}}}
	fmt.Println(getIntersectionNode(a, b) == nil)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
