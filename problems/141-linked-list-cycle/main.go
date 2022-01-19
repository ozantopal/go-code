package main

import (
	"fmt"

	"github.com/ozantopal/go-code/structures"
)

type ListNode = structures.ListNode

func main() {
	sample1()
	sample2()
	sample3()
}

func sample1() {
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n1.Next = n2
	n3 := &ListNode{Val: 0}
	n2.Next = n3
	n4 := &ListNode{Val: -4}
	n3.Next = n4
	n4.Next = n2
	fmt.Println(hasCycle(n1))
}

func sample2() {
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n1.Next = n2
	n2.Next = n1
	fmt.Println(hasCycle(n1))
}

func sample3() {
	n1 := &ListNode{Val: 3}
	fmt.Println(hasCycle(n1))
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}
