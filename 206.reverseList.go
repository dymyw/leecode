package dymyw_leecode

// tags: LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 多指针
	var prev *ListNode
	cur := head

	for cur != nil {
		// 后移, 后移, 反转
		prev, cur, cur.Next = cur, cur.Next, prev
	}

	// cur == nil
	return prev
}
