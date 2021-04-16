package LinkedList

// DoublyListNode 双向链表节点
type DoublyListNode struct {
	Val int
	Next *DoublyListNode
	Prev *DoublyListNode
}

// DoublyLinkedList 双向链表
type DoublyLinkedList struct {
	head *DoublyListNode
	tail *DoublyListNode
	size uint
}
