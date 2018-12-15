package kit

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListToInts(head *ListNode) []int {
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("Limit exceeded %d", limit)
			panic(msg)
		}
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
