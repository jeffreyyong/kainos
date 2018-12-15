package problem0019

// ListNode is a node
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	parent, headIsNthFromEnd := getParent(head, n)

	if headIsNthFromEnd {
		// Delete the head node
		return head.Next
	}
	parent.Next = parent.Next.Next

	return head
}

func getParent(head *ListNode, n int) (parent *ListNode, headIsNthFromEnd bool) {
	parent = head

	for head != nil {
		if n < 0 {
			parent = parent.Next
		}
		n--
		head = head.Next
	}

	headIsNthFromEnd = n == 0
	return
}
