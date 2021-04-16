package Stack

// Stacker 栈
type Stacker interface {
	// Push 入栈
	Push(value interface{})

	// Pop 出栈
	Pop() interface{}
}
