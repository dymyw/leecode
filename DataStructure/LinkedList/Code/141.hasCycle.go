package Code

// 数据结构：LinkedList
// 算法：多指针

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// 定义指针
	fast := head

	// 循环条件，递归的指针
	for fast != nil && fast.Next != nil {
		// 联动
		head = head.Next
		fast = fast.Next.Next

		// 值计算
		if fast == head {
			return true
		}
	}

	// 返回计算的结果
	return false
}
