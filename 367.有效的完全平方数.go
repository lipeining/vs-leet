/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num <= 1 {
		return true
	}
	l := 0
	r := num
	for l < r {
		mid := l+(r-l)/2
		m := mid * mid
		if m == num {
			return true
		} else if m > num {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return false
}
// @lc code=end

