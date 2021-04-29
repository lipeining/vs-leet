/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	num := x^y
	ans := 0
	for num != 0 {
		num &= num-1
		ans++
	}
	return ans
}
// @lc code=end

