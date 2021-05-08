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
func (as *ArrayStack) Push(value interface{}) {
	// 追加数组元素
	as.arr = append(as.arr, value)
}

// Pop 出栈
func (as *ArrayStack) Pop() interface{} {
	// 判断是否为空
	if as.Empty() {
		return nil
	}

	// 数据截断，获取最后一个元素
	lastPos := len(as.arr) - 1
	x := as.arr[lastPos]
	as.arr = as.arr[:lastPos]

	return x
}

// Top 获取栈顶元素
func (as *ArrayStack) Top() interface{} {
	// 判断是否为空
	if as.Empty() {
		return nil
	}

	lastPos := len(as.arr) - 1

	return as.arr[lastPos]
}

// Empty 是否为空
func (as *ArrayStack) Empty() bool {
	return len(as.arr) == 0
}
