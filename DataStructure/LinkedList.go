package DataStructure

// ListNode 链表节点
type ListNode struct {
	Val int
	Next *ListNode
}

// DoublyListNode 链表节点
type DoublyListNode struct {
	Val int
	Next *DoublyListNode
	Prev *DoublyListNode
}

// LinkedList 单链表
type LinkedList struct {
	head *ListNode
	size uint
}

// DoublyLinkedList 双向链表
type DoublyLinkedList struct {
	head *ListNode
	tail *ListNode
	size uint
}
