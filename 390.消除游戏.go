/*
 * @lc app=leetcode.cn id=390 lang=golang
 *
 * [390] 消除游戏
 */

// @lc code=start
func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}
	return 2 * (n/2 + 1 - lastRemaining(n/2))
}

// @lc code=end

