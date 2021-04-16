package Queue

import "dymyw-leecode/DataStructure/LinkedList"

// DoublyLinkedListQueue 双向链表实现的队列
type DoublyLinkedListQueue struct {
	list *LinkedList.DoublyLinkedList
}

// NewDoublyLinkedListQueue 新建一个双向链表队列
func NewDoublyLinkedListQueue() *DoublyLinkedListQueue {
	return &DoublyLinkedListQueue{}
}

// Push 入队
func (dllq *DoublyLinkedListQueue) Push(value interface{}) {
	list := dllq.list

	node := &LinkedList.DoublyListNode{
		Val: value,
		Prev: nil,
		Next: list.Head,
	}
	list.Head = node
	list.Size += 1
}

// Pop 出队
func (dllq *DoublyLinkedListQueue) Pop() interface{} {
	if dllq.Empty() {
		return nil
	}

	list := dllq.list
	node := list.Tail
	list.Tail = list.Tail.Prev
	list.Size -= 1

	return node.Val
}

// Peek 获取最前的元素
func (dllq *DoublyLinkedListQueue) Peek() interface{} {
	if dllq.Empty() {
		return nil
	}

	return dllq.list.Head.Val
}

// Empty 是否为空
func (dllq *DoublyLinkedListQueue) Empty() bool {
	return dllq.list.Size == 0
}
