/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H 指数 II
 */

// @lc code=start
func hIndex(citations []int) int {
	n := len(citations)
	for idx, c := range citations {
		if c >= n-idx {
			return n - idx
		}
	}
	return 0
}

// @lc code=end

