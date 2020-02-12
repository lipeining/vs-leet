/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	return max(robStartEnd(nums, 0, length-2), robStartEnd(nums, 1, length-1))
}
func robStartEnd(nums []int, start int, end int) int {
	prev := 0
	next := 0
	for i := start; i <= end; i++ {
		now := max(prev+nums[i], next)
		prev = next
		next = now
	}
	return next
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
