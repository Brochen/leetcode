package main

import "fmt"

/**

Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

 */

type ListNode struct {
     Val int
     Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var head, p *ListNode

	m := 0
	for l1 != nil || l2 != nil {
		n1 := 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		n2 := 0
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		n := n1 + n2 + m

		if n >= 10 {
			n -= 10
			m = 1
		} else {
			m = 0
		}

		pt := &ListNode{n, nil}
		if p != nil {
			p.Next = pt
		}
		p = pt

		if head == nil {
			head = p
		}

	}
	if m > 0 {
		pt := &ListNode{m, nil}
		p.Next = pt
	}

	return head
}

func main()  {
	fmt.Println("====111")
}
