/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(N int) int {
    if N <= 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}
// @lc code=end

