package LinkedList

// ListNode 链表节点
type ListNode struct {
	Val interface{}
	Next *ListNode
}

// LinkedList 单链表
type LinkedList struct {
	Head *ListNode
	Size uint
}
