/*
 * @lc app=leetcode.cn id=1438 lang=golang
 *
 * [1438] 绝对差不超过限制的最长连续子数组
 */

// @lc code=start
func longestSubarray(nums []int, limit int) int {
	n := len(nums)
	max := make([]int, 0)
	min := make([]int, 0)
	l, r := 0, 0
	ans := 0
	for r < n {
		for len(max) != 0 && max[len(max)-1] < nums[r] {
			max = max[:len(max)-1]
		}
		for len(min) != 0 && min[len(min)-1] > nums[r] {
			min = min[:len(min)-1]
		}
		max = append(max, nums[r])
		min = append(min, nums[r])
		for max[0]-min[0] > limit {
			if max[0] == nums[l] {
				max = max[1:]
			}
			if min[0] == nums[l] {
				min = min[1:]
			}
			l++
		}
		now := r - l + 1
		if now > ans {
			ans = now
		}
		r++
	}
	return ans
}

// @lc code=end

