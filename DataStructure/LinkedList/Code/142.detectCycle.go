package Code

// 数据结构：LinkedList
// 算法：记忆化搜索
	// todo 双指针：快慢指针 O(1) 的空间复杂度

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// 引入 set
	set := make(map[*ListNode]bool)

	// 循环条件
	for head != nil {
		// 记忆化搜索 O(n)
		if _, exists := set[head]; exists {
			return head
		}

		// 记忆
		set[head] = true

		// 后移
		head = head.Next
	}

	// 返回结果
	return nil
}
