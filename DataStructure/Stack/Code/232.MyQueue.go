package Code

// 数据结构：Stack、Array
// 算法：数组栈基本操作; outStack 为空，把 inStack 中的成员迁移到 outStack;

type MyQueue struct {
	inStack, outStack []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	// 无脑入
	this.inStack = append(this.inStack, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	// outStack 为空，把 inStack 中的成员迁移到 outStack
	if len(this.outStack) == 0 {
		this.in2out()
	}

	// 无脑出
	lastOutPos := len(this.outStack) - 1
	x := this.outStack[lastOutPos]
	this.outStack = this.outStack[:lastOutPos]

	return x
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.in2out()
	}

	lastOutPos := len(this.outStack) - 1
	return this.outStack[lastOutPos]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}


// outStack.Push(inStack.Pop())
func (this *MyQueue) in2out() {
	for len(this.inStack) > 0 {
		lastInPos := len(this.inStack) - 1
		x := this.inStack[lastInPos]
		this.inStack = this.inStack[:lastInPos]

		this.outStack = append(this.outStack, x)
	}
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
