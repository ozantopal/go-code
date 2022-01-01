package main

import "fmt"

func main() {
	l1, l2 := ListNode{5,nil}, ListNode{7,nil}
	result := addTwoNumbers(&l1, &l2)

	fmt.Print(result.Next.Val)
	fmt.Print(result.Val)
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry, head := 0, l1
	for {
		l1.Val += l2.Val + carry
		carry = int(l1.Val / 10)
		l1.Val = l1.Val % 10

		if l2.Next == nil {
			break
		} else if l1.Next == nil {
			l1.Next = l2.Next
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	for carry != 0 {
        if (l1.Next == nil) { 
			l1.Next = &ListNode{0, nil} 
		}
        
        l1.Next.Val += carry        
        carry = l1.Next.Val / 10
        l1.Next.Val = l1.Next.Val % 10
        
        l1 = l1.Next
    }
	
	return head
}

type ListNode struct {
    Val int
    Next *ListNode
}