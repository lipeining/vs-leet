/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	n := len(nums)
	to := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		if i <= to {
			to = max(to, nums[i]+i)
			if to >= n-1 {
				return true
			}
		}
	}
	return false

	// dp
	// if len(nums) <= 1 {
	// 	return true
	// }
	// dp := make([]bool, len(nums))
	// dp[0] = true
	// for i := 0; i < len(nums); i++ {
	// 	if !dp[i] {
	// 		return false
	// 	}
	// 	step := nums[i]
	// 	// if i+step >= len(nums) {
	// 	// 	step = len(nums) - i - 1
	// 	// }
	// 	for j := 1; j <= step; j++ {
	// 		if i+j >= len(nums) {
	// 			break
	// 		}
	// 		dp[i+j] = true
	// 	}
	// }
	// return dp[len(nums)-1]
}
func dfs(nums []int, now int) bool {
	if now >= len(nums)-1 {
		return true
	}
	max := nums[now] + now
	if max > len(nums) {
		max = len(nums)
	}
	for i := now + 1; i <= max; i++ {
		if dfs(nums, i) {
			return true
		}
	}
	return false
}

// @lc code=end

