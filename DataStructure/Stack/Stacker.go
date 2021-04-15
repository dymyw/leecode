package Stack

// 栈
type Stacker interface {
	// 入栈
	Push(value interface{})

	// 出栈
	Pop() interface{}
}
