/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	n := len(intervals)
	ans := make([][]int, 0)
	merge := false
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		a, b := intervals[i][0], intervals[i][1]
		if a > right {
			if !merge {
				ans = append(ans, []int{left, right})
				merge = true
			}
			ans = append(ans, intervals[i])
		} else if b < left {
			// do nothing
			ans = append(ans, intervals[i])
		} else {
			left = min(left, a)
			right = max(right, b)
		}
	}
	if !merge {
		ans = append(ans, []int{left, right})
	}
	return ans
}

// @lc code=end

