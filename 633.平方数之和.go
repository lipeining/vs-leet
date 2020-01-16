/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */

// @lc code=start
func judgeSquareSum(c int) bool {
	for a:= 0; a * a <= c; a++ {
		b := math.Sqrt(float64(c - a * a))
		n := float64(int(b))
		if n == b {
			return true
		}
	}
	return false
}
// @lc code=end

