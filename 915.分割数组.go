/*
 * @lc app=leetcode.cn id=915 lang=golang
 *
 * [915] 分割数组
 */

// @lc code=start
func partitionDisjoint(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n := len(nums)
	// maxL := make([]int, n)
	// maxL[0] = nums[0]
	// for i := 1; i < n; i++ {
	// 	maxL[i] = max(nums[i], maxL[i-1])
	// }
	minR := make([]int, n)
	minR[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		minR[i] = min(minR[i+1], nums[i])
	}
	left := nums[0]
	for i := 1; i < n; i++ {
		if left <= minR[i] {
			return i
		}
		left = max(left, nums[i])
	}
	return -1
}

// @lc code=end

