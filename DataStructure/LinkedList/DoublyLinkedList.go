package LinkedList

// DoublyListNode 双向链表节点
type DoublyListNode struct {
	Val interface{}
	Next *DoublyListNode
	Prev *DoublyListNode
}

// DoublyLinkedList 双向链表
type DoublyLinkedList struct {
	Head *DoublyListNode
	Tail *DoublyListNode
	Size uint
}
