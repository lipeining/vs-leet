/*
 * @lc app=leetcode.cn id=895 lang=golang
 *
 * [895] 最大频率栈
 */

// @lc code=start
type FreqStack struct {
	freq map[int]int
	group map[int][]int
	maxf int
}


func Constructor() FreqStack {
	return FreqStack{
		freq: make(map[int]int),
		group: make(map[int][]int),
		maxf: 0,
	}
}


func (this *FreqStack) Push(val int)  {
	f := this.freq[val] + 1
	this.freq[val] = f
	if f > this.maxf {
		this.maxf = f
	}
	this.group[f] = append(this.group[f], val)
}


func (this *FreqStack) Pop() int {
	size := len(this.group[this.maxf])
	x := this.group[this.maxf][size-1]
	this.group[this.maxf] = this.group[this.maxf][:size-1]
	this.freq[x]--
	if len(this.group[this.maxf]) == 0 {
		this.maxf--
	}
	return x
}


/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
// @lc code=end

