/*
 * @lc app=leetcode.cn id=495 lang=golang
 *
 * [495] 提莫攻击
 */

// @lc code=start
func findPoisonedDuration(timeSeries []int, duration int) int {
	n := len(timeSeries)
	if n == 0 {
		return 0
	}
	ans := 0
	min := func(a, b int)int{
		if a < b {
			return a
		}
		return b
	}
	for i := 0;i < n-1; i++ {
		ans += min(timeSeries[i+1]-timeSeries[i], duration)
	}
	return ans + duration
}
// @lc code=end

