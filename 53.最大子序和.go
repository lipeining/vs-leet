/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	dp_i := nums[0]
	ans := nums[0]
	for i := 1; i < length; i++ {
		dp_i = max(nums[i], dp_i+nums[i])
		ans = max(ans, dp_i)
	}
	return ans
	// length := len(nums)
	// if length == 0 {
	// 	return 0
	// }
	// dp := make([]int, length)
	// // dp 指以 i 为结尾的和最大
	// // 如果当前
	// // dp[i] = max(nums[i], dp[i-1]+nums[i])
	// dp[0] = nums[0]
	// for i := 1; i < length; i++ {
	// 	dp[i] = max(nums[i], dp[i-1]+nums[i])
	// }
	// fmt.Println(dp)
	// max := dp[0]
	// for i := 0; i < length; i++ {
	// 	if max < dp[i] {
	// 		max = dp[i]
	// 	}
	// }
	// return max

	// 贪心算法
	// length := len(nums)
	// if length == 0 {
	// 	return 0
	// }
	// // dp 指以 i 为结尾的和最大
	// // 如果当前
	// // dp[i] = max(nums[i], dp[i-1]+nums[i])
	// curSum := nums[0]
	// maxSum := nums[0]
	// for i := 1; i < length; i++ {
	// 	curSum = max(nums[i], curSum+nums[i])
	// 	maxSum = max(curSum, maxSum)
	// }
	// return maxSum
	// length := len(nums)
	// if length == 0 {
	// 	return 0
	// }
	// dp := make([]int, length)
	// // dp 指以 i 为结尾的和最大
	// // 如果当前
	// // dp[i] = max(nums[i], dp[i-1]+nums[i])
	// dp[0] = nums[0]
	// for i := 1; i < length; i++ {
	// 	dp[i] = max(nums[i], dp[i-1]+nums[i])
	// }
	// fmt.Println(dp)
	// max := dp[0]
	// for i := 0; i < length; i++ {
	// 	if max < dp[i] {
	// 		max = dp[i]
	// 	}
	// }
	// return max
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
