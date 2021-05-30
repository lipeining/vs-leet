/*
 * @lc app=leetcode.cn id=991 lang=golang
 *
 * [991] 坏了的计算器
 */

// @lc code=start
func brokenCalc(x int, y int) int {
	if y <= x {
		return x - y
	}
	if y%2 == 0 {
		return 1 + brokenCalc(x, y/2)
	}
	return 1 + brokenCalc(x, y+1)
}

// @lc code=end

