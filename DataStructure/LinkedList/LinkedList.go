package LinkedList

// ListNode 链表节点
type ListNode struct {
	Val int
	Next *ListNode
}

// LinkedList 单链表
type LinkedList struct {
	head *ListNode
	size uint
}
