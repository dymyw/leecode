package Queue

// ArrayQueue 数组实现的队列
type ArrayQueue struct {
	arr []interface{}
}

// NewArrayQueue 新建一个数组队列
func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{}
}

// Push 入队
func (aq *ArrayQueue) Push(value interface{}) {
	aq.arr = append(aq.arr, value)
}

// Pop 出队
func (aq *ArrayQueue) Pop() interface{} {
	if aq.Empty() {
		return nil
	}

	x := aq.arr[0]
	aq.arr = aq.arr[1:]

	return x
}

// Peek 获取最前的元素
func (aq *ArrayQueue) Peek() interface{} {
	if aq.Empty() {
		return nil
	}

	return aq.arr[0]
}

// Empty 是否为空
func (aq *ArrayQueue) Empty() bool {
	return len(aq.arr) == 0
}
