package Stack

// 数组实现的栈
type ArrayStack struct {
	arr []interface{}
}

// 新建一个数组栈
func NewArrayStack() *ArrayStack {
	return &ArrayStack{}
}

// 入栈
func (as *ArrayStack) Push(x interface{}) {
	as.arr = append(as.arr, x)
}

// 出栈
func (as *ArrayStack) Pop() interface{} {
	lastPos := len(as.arr) - 1
	x := as.arr[lastPos]
	as.arr = as.arr[:lastPos]

	return x
}
