package main

import (
	"fmt"
)

// ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// 非递归方式,从头到尾移动节点进行相加并记录进位值
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0}
	curr := dummyHead

	// 进位
	carry := 0

	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10

		// 如果sum是两位数,则取个位数
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next

}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l3 := &ListNode{
		Val: 5,
	}

	l4 := &ListNode{
		Val: 5,
	}

	fmt.Println(addTwoNumbers(l1, l2))
	fmt.Println(addTwoNumbers(l3, l4))
}
