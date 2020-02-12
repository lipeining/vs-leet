/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	// 和零钱不同的是，顺序不同就是不一样的结果
	dp := make([]int, target+1)
	// dp[0] = 1
	// for i := 0; i < len(nums); i++ {
	// 	num := nums[i]
	// 	for j := num; j <= target; j++ {
	// 		dp[j] += dp[j-num]
	// 	}
	// }
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i-nums[j] >= 0 {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}

// @lc code=end
