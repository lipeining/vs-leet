/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
    for n != 0 {
		m := n % 2
		n = n /2 
		if m == 1 && n != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

