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
	// 引入集合
	set := map[*ListNode]bool{}

	for head != nil {
		// 记忆化搜索 O(n)
		if _, exists := set[head]; exists {
			return true
		} else {
			// 记忆
			set[head] = true
			// 后移
			head = head.Next
		}
	}

	return false
}
