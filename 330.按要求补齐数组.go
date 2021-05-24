/*
 * @lc app=leetcode.cn id=330 lang=golang
 *
 * [330] 按要求补齐数组
 */

// @lc code=start
func minPatches(nums []int, n int) int {
	ans := 0
	for i, x := 0, 1; x <= n; {
		if i < len(nums) && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			x *= 2
			ans++
		}
	}
	return ans
}

// @lc code=end

