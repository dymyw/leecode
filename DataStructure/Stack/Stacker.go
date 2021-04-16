package Stack

// Stacker 栈
type Stacker interface {
	// Push 入栈
	Push(value interface{})

	// Pop 出栈
	Pop() interface{}

	// Top 获取栈顶元素
	Top() interface{}

	// Empty 是否为空
	Empty() bool
}
