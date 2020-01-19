/*
 * @lc app=leetcode.cn id=933 lang=golang
 *
 * [933] 最近的请求次数
 */

// @lc code=start
type RecentCounter struct {
    queue []int
}


func Constructor() RecentCounter {
    return RecentCounter{}
}


func (this *RecentCounter) Ping(t int) int {
	// t-3000, t
	left := t-3000
	// right := t
	i:=0;
	this.queue = append(this.queue, t)
	for i<len(this.queue) {
		if this.queue[i] >= left {
			break
		}
		i++
	}
	this.queue = this.queue[i:]
	return len(this.queue)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// @lc code=end

