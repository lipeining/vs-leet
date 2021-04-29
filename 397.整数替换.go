/*
 * @lc app=leetcode.cn id=397 lang=golang
 *
 * [397] 整数替换
 */

// @lc code=start
func integerReplacement(n int) int {
	if n == math.MaxInt32 {
		return 32
	}
	if n <= 3 {
		return n - 1
	}
	if n%2 == 0 {
		return integerReplacement(n/2) + 1
	}
	if n%2 == 0 {
		return integerReplacement(n-1) + 1
	}
	return integerReplacement(n+1) + 1
}

// @lc code=end

