/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H指数 II
 */

// @lc code=start
func hIndex(citations []int) int {
	idx := 0
	n := len(citations)
	for i := 0; i < n; i++ {
		if citations[i] >= n-idx {
			return n - idx
		} else {
			idx++
		}
	}
	return 0
}

// @lc code=end
