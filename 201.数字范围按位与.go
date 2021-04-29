/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] 数字范围按位与
 */

// @lc code=start
func rangeBitwiseAnd(m int, n int) int {
	for m < n {
		n &= n-1
	}
	return n
	// shift := 0
	// for m < n {
	// 	m >>= 1
	// 	n >>= 1
	// 	shift++
	// }
	// return 1 << shift
}

// @lc code=end

