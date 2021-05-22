/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	in  []int
	out []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		for len(this.in) != 0 {
			this.out = append(this.out, this.in[len(this.in)-1])
			this.in = this.in[:len(this.in)-1]
		}
	}
	v := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return v
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		for len(this.in) != 0 {
			this.out = append(this.out, this.in[len(this.in)-1])
			this.in = this.in[:len(this.in)-1]
		}
	}
	v := this.out[len(this.out)-1]
	return v
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

