package main

import (
	"fmt"

	"github.com/ozantopal/go-code/structures"
)

type ListNode = structures.ListNode

func main() {
	h1 := structures.Ints2List([]int{1, 2, 6, 3, 4, 5, 6})
	fmt.Println(structures.List2Ints(removeElements(h1, 6)))

	h2 := structures.Ints2List([]int{})
	fmt.Println(structures.List2Ints(removeElements(h2, 1)))

	h3 := structures.Ints2List([]int{7, 7, 7, 7})
	fmt.Println(structures.List2Ints(removeElements(h3, 7)))
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{Next: head}
	pre := newHead
	curr := head

	for curr != nil {
		if curr.Val == val {
			pre.Next = curr.Next
		} else {
			pre = curr
		}

		curr = curr.Next
	}

	return newHead.Next
}
