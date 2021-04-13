package dymyw_leecode

// tags: LinkedList、Set

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// 引入集合
	set := map[*ListNode]bool{}

	for head != nil {
		// 记忆化搜索 O(n)
		if _, exists := set[head]; exists {
			return head
		} else {
			// 记忆
			set[head] = true
			// 后移
			head = head.Next
		}
	}

	return nil
}
