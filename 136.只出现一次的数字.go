/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	a := 0
	for _, num := range nums {
		a ^= num
	}
	return a ^ 0
}

// @lc code=end

