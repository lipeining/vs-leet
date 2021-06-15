/*
 * @lc app=leetcode.cn id=352 lang=golang
 *
 * [352] 将数据流变为多个不相交区间
 */

// @lc code=start
type SummaryRanges struct {
	inter [][]int
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{inter: make([][]int, 0)}
}

func (this *SummaryRanges) AddNum(val int) {
	n := len(this.inter)
	l, r := 0, n-1
	for l < r {
		mid := (l + r) >> 1
		set := this.inter[mid]
		if set[0] <= val && set[1] >= val {
			// in interval
			return
		} else if set[0] > val {
			r = mid - 1
		} else if set[1] < val {
			l = mid + 1
		} else {
			// never happen
		}
	}
	if l == n-1 {
		// insert
		this.inter = append(this.inter, []int{val, val})
		return
	}
	if l < 0 {
		this.inter = append([]int{val, val}, this.inter...)
		return
	}
	set = this.inter[l]
	if set[0] > val {

	}
	if set[0] < val {

	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.inter
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
// @lc code=end

