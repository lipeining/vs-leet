/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	n := len(nums)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		if dp[i] == math.MaxInt32 {
			continue
		}
		dis := nums[i]
		for j := 1; j <= dis; j++ {
			if i+j < n {
				dp[j+i] = min(dp[j+i], dp[i]+1)
			}
		}
	}
	return dp[n-1]
}

// @lc code=end

