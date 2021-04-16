package Queue

// Queuer 队列
type Queuer interface {
	// Push 入队
	Push(value interface{})

	// Pop 出队
	Pop() interface{}

	// Peek 获取最前的元素
	Peek() interface{}

	// Empty 是否为空
	Empty() bool
}
