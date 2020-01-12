/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(num int) bool {
	n:=num
	if n < 1 {
		return false
	}
	for n % 4==0 {
		n = n/4
	}
	return n==1
}
// @lc code=end

