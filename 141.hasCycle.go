package dymyw_leecode

// tags: LinkedList、Set

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// 多指针
	fast := head

	// 循环条件
	for fast != nil && fast.Next != nil {
		// 4. 联动
		head = head.Next
		fast = fast.Next.Next
		// 计算
		if fast == head {
			return true
		}
	}

	return false
}
