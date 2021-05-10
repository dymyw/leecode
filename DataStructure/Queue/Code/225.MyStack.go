package Code

// 数据结构：Queue、Array
// 算法：数组队列基本操作; 元素入队后，其它的元素重新依次出队，再入队;

type MyStack struct {
	inQueue, outQueue []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	// 其实用一个环就可以实现，把一个队列看成一个环

	// 把元素入队
	this.inQueue = append(this.inQueue, x)

	// 其它的元素重新依次出队，再入队
	for len(this.outQueue) > 0 {
		p := this.outQueue[0]
		this.outQueue = this.outQueue[1:]

		this.inQueue = append(this.inQueue, p)
	}

	// 把空的队列设为 inQueue
	this.inQueue, this.outQueue = this.outQueue, this.inQueue
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	x := this.outQueue[0]
	this.outQueue = this.outQueue[1:]

	return x
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.outQueue[0]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.outQueue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
