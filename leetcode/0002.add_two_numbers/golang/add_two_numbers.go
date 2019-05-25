package main

import "fmt"

func main() {

	// group 1
	listNode13 := &ListNode{Val: 3, Next: nil}
	listNode12 := &ListNode{Val: 4, Next: listNode13}
	listNode11 := &ListNode{Val: 2, Next: listNode12}

	// group 2
	listNode23 := &ListNode{Val: 4, Next: nil}
	listNode22 := &ListNode{Val: 6, Next: listNode23}
	listNode21 := &ListNode{Val: 5, Next: listNode22}

	listNodeAnswer := addTwoNumbers(listNode11, listNode21)

	for {
		// head := new(ListNode)
		// cur := head
		if listNodeAnswer != nil {
			number1 := listNodeAnswer.Val
			listNodeAnswer = listNodeAnswer.Next
			fmt.Println(number1)
		} else {
			return
		}
	}
}

// ListNode contains a Val and the Next ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	cur := head
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}

		num := n1 + n2 + carry
		carry = num / 10
		cur.Next = &ListNode{Val: num % 10, Next: nil}
		cur = cur.Next
	}
	return head.Next
}
