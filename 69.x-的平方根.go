/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	l := 0
	r := x
	for l < r {
		// 取右边中位数
		mid := l + (r-l+1)/2
		now := mid * mid
		if now > x {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

// @lc code=end

