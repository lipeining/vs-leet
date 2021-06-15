/*
 * @lc app=leetcode.cn id=731 lang=golang
 *
 * [731] 我的日程安排表 II
 */

// @lc code=start
type MyCalendarTwo struct {
	calendar [][]int
	overlap [][]int
}


func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		calendar: make([][]int, 0),
		overlap: make([][]int, 0),
	}
}

func max(a, b int)int{
	if a > b {
		return a
	}
	return b
}

func min(a, b int)int{
	if a < b {
		return a
	}
	return b
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, iv := range this.overlap {
		if iv[0] < end && iv[1] > start {
			return false
		}
	}
	for _, iv := range this.calendar {
		if iv[0] < end && iv[1] > start {
			this.overlap = append(this.overlap, []int{max(start, iv[0]), min(end, iv[1])})
		}
	}
	this.calendar = append(this.calendar, []int{start, end})
	return true
}


/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end

