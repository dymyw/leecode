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

	// 循环
	for cur != nil {
		//// 计算
		//x := cur.Next
		//cur.Next = prev
		//// 4. 联动
		//prev, cur = cur, x

		// 反转，后移，后移
		cur.Next, prev, cur = prev, cur, cur.Next
	}

	// cur == nil
	return prev
}
