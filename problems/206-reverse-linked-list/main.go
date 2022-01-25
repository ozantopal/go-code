package main

import (
	"fmt"

	"github.com/ozantopal/go-code/structures"
)

type ListNode = structures.ListNode

func main() {
	head := structures.Ints2List([]int{1, 2, 3, 4, 5})
	reversedArr := structures.List2Ints(reverseList(head))

	fmt.Println(reversedArr)
}

func reverseList(head *ListNode) *ListNode {
	var tail *ListNode
	for head != nil {
		next := head.Next
		head.Next = tail
		tail = head
		head = next
	}
	return tail
}
