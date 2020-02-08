import "fmt"

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	// 以 i 结尾的最长上升子序列
	// dp[i] = max(dp[i-1] + 1, 1) if nums[i] > nums[i-1]
	// 不是连续的，需要考虑 跳步行为

	length := len(nums)
	if length == 0 {
		return 0
	}
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = 1
	}
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	fmt.Println(dp)
	// return dp[length-1]
	ans := dp[0]
	for i := 1; i < length; i++ {
		if ans < dp[i] {
			ans = dp[i]
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
