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
func reverseList(head *ListNode) *ListNode {
	// 定义指针
	var prev *ListNode
	cur := head

	// 循环条件，递归的指针
	for cur != nil {
		// // 指针运算
		// temp := cur.Next
		// cur.Next = prev

		// // 联动
		// prev = cur
		// cur = temp

		// 简化
		cur.Next, prev, cur = prev, cur, cur.Next
	}

	// 返回某个指针
	return prev
}
