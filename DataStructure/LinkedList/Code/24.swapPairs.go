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
func swapPairs(head *ListNode) *ListNode {
	// 定义指针
	node := &ListNode{}
	node.Next = head
	prev := node

	// 循环条件，递归的指针
	for head != nil && head.Next != nil {
		// 指针运算
		next := head.Next
		head.Next = next.Next
		next.Next = head
		prev.Next = next

		// 联动
		prev = head
		head = head.Next
	}

	// 返回某个指针
	return node.Next
}
