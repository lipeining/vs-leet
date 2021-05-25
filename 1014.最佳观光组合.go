/*
 * @lc app=leetcode.cn id=1014 lang=golang
 *
 * [1014] 最佳观光组合
 */

// @lc code=start
func maxScoreSightseeingPair(values []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans, mx := 0, 0
	for i := 0; i < len(values); i++ {
		ans = max(ans, values[i]-i+mx)
		mx = max(mx, values[i]+i)
	}
	return ans
}

// @lc code=end

