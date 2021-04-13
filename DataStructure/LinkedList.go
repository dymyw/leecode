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
