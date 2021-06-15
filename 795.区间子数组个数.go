/*
 * @lc app=leetcode.cn id=795 lang=golang
 *
 * [795] 区间子数组个数
 */

// @lc code=start
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	// max := func(a, b int) int {
	// 	if a > b {
	// 		return a
	// 	}
	// 	return b
	// }
	c := 0
	b := -1
	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > right {
			b = i
			c = 0
		} else if nums[i] < left {
			ans += c
		} else {
			c = i - b
			ans += c
		}
	}
	return ans
}

// @lc code=end

