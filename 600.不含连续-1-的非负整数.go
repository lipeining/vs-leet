/*
 * @lc app=leetcode.cn id=600 lang=golang
 *
 * [600] 不含连续1的非负整数
 */

// @lc code=start
func findIntegers(num int) int {
	var find func(i, sum, num int, prev bool) int
	find = func(i, sum, num int, prev bool) int {
		if sum > num {
			return 0
		}
		if 1<<i > num {
			return 1
		}
		if prev {
			// 前一个是1
			return find(i+1, sum, num, false)
		}
		return find(i+1, sum, num, false) + find(i+1, sum+(1<<i), num, true)
	}
	return find(0, 0, num, false)
}

// @lc code=end

