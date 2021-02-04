/*
 * @lc app=leetcode.cn id=1186 lang=golang
 *
 * [1186] 删除一次得到子数组最大和
 */

// @lc code=start
func maximumSum(arr []int) int {
	n := len(arr)
	// 他人答案
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = arr[0]
	dp[0][1] = math.MinInt32
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := max(dp[0][0], dp[0][1])
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0]+arr[i], arr[i])
		left := dp[i-1][1]
		if left == math.MinInt32 {
			left = 0
		}
		// 可能删除当前arr[i]
		dp[i][1] = max(left+arr[i], dp[i-1][0])
		ans = max(ans, max(dp[i][0], dp[i][1]))
	}
	return ans
}

// @lc code=end

