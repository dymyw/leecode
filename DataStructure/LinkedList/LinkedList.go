package LinkedList

// 数据结构：LinkedList

// 特点：
//	利用零散的内存空间

// 时间复杂度：
//	Access：O(n)
//	Insert：平均 O(1)
//	Delete：平均 O(1)

// 优点：
//	改善插入、删除

// 适用场景：
//	不确定元素的个数时

// 特殊情况：
//	双向链表
//	环状链表

// head := &ListNode{0, nil}
// var prev *ListNode

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
