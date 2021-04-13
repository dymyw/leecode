package dymyw_leecode

// tags: LinkedList

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 新建一个新的链表
	head := &ListNode{}
	// 指针赋值操作
	cur := head
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		// 计算: 相加
		sum := carry
		if l1 != nil {
			sum += l1.Val
			// 后移
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			// 后移
			l2 = l2.Next
		}

		carry = sum / 10
		// 指针赋值操作
		cur.Next = &ListNode{sum % 10, nil}

		// 后移
		cur = cur.Next
	}

	return head.Next
}
