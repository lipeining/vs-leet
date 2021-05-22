/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	q []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		q: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	n := len(this.q)
	this.q = append(this.q, x)
	for ; n > 0; n-- {
		this.q = append(this.q, this.q[0])
		this.q = this.q[1:]
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	v := this.q[0]
	this.q = this.q[1:]
	return v
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

