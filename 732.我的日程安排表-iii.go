/*
 * @lc app=leetcode.cn id=732 lang=golang
 *
 * [732] 我的日程安排表 III
 */

// @lc code=start
type MyCalendarThree struct {
	delta map[int]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{delta: make(map[int]int)}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	this.delta[start]++
	this.delta[end]--
	active := 0
	ans := 0
	keys := make([]int, 0)
	for k := range this.delta {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		active += this.delta[k]
		if active > ans {
			ans = active
		}
	}
	return ans
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end

