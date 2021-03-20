/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	ret := 0
	for _,num := range nums {
		ret ^= num
	}
	div := 1
	for div & ret == 0 {
		div <<= 1
	}
	a, b := 0, 0
	for _, n := range nums {
		if div & n > 0 {
			a ^= n
		} else {
			b ^= n
		}
	}
	return []int{a, b}
}
// @lc code=end

