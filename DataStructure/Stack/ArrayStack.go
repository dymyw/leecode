package Stack

// ArrayStack 数组实现的栈
type ArrayStack struct {
	arr []interface{}
}

// NewArrayStack 新建一个数组栈
func NewArrayStack() *ArrayStack {
	return &ArrayStack{}
}

// Push 入栈
func (as *ArrayStack) Push(x interface{}) {
	as.arr = append(as.arr, x)
}

// Pop 出栈
func (as *ArrayStack) Pop() interface{} {
	lastPos := len(as.arr) - 1
	x := as.arr[lastPos]
	as.arr = as.arr[:lastPos]

	return x
}
