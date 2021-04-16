package Stack

import "dymyw-leecode/DataStructure/LinkedList"

// LinkedListStack 单链表实现的栈
type LinkedListStack struct {
	node *LinkedList.ListNode
}

// NewLinkedListStack 新建一个链表栈
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{}
}

// Push 入栈
func (lls *LinkedListStack) Push(value interface{}) {
	node := &LinkedList.ListNode{
		Val: value,
		Next: lls.node,
	}
	lls.node = node
}

// Pop 出栈
func (lls *LinkedListStack) Pop() interface{} {
	x := lls.node.Val
	lls.node = lls.node.Next

	return x
}

// Top 获取栈顶元素
func (lls *LinkedListStack) Top() interface{} {
	return lls.node.Val
}

// Empty 是否为空
func (lls *LinkedListStack) Empty() bool {
	return lls.node == nil
}

