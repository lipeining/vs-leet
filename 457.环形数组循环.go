/*
 * @lc app=leetcode.cn id=457 lang=golang
 *
 * [457] 环形数组循环
 */

// @lc code=start
func circularArrayLoop(nums []int) bool {
	length := len(nums)
	for start := 0; start < length; start++ {
		if loop(nums, start) {
			return true
		}
	}
	return false
}
func loop(nums []int, start int) bool {
	len := 0
	length := len(nums)
	curr := start
	return false
}

// @lc code=end
