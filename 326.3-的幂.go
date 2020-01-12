/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 */

// @lc code=start
func isPowerOfThree(n int) bool {
    if n < 1 {
		return false
	}
	for n % 3==0 {
		n = n/3
	}
	return n==1
}
// @lc code=end

