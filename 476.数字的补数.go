/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */

// @lc code=start
func findComplement(num int) int {
	res := 0
	multi := 1
    for num != 0 {
		m := num% 2
		num = num / 2
		if m == 0 {
			res += multi
		}
		multi = multi * 2
	}
	return res
}
// @lc code=end

