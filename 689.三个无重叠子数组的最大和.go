/*
 * @lc app=leetcode.cn id=689 lang=golang
 *
 * [689] 三个无重叠子数组的最大和
 */

// @lc code=start
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	return maxSumOfNSubarrays(nums, k, 3)
}
func maxSumOfNSubarrays(nums []int, k int, n int) []int {
	N := len(nums)
	if k < 1 || n*k > N {
		return []int{}
	}
	s := 0
	for i := 0; i < k; i++ {
		s += nums[i]
	}
	sums := make([]int, N)
	sums[k-1] = s
	for i := k; i < N; i++ {
		s += nums[i] - nums[i-k]
		sums[i] = s
	}
	dp := make([][]int, N)
	path := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, n+1)
		path[i] = make([]int, n+1)
	}
	dp[k-1][1] = sums[k-1]
	path[k-1][1] = k - 1
	for i := k; i < N; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j]
			path[i][j] = path[i-1][j]
			if sums[i]+dp[i-k][j-1] > dp[i][j] {
				dp[i][j] = sums[i] + dp[i-k][j-1]
				path[i][j] = i
			}
		}
	}
	res := make([]int, 0)
	// 路径回溯
	ind := path[N-1][n]
	res = append(res, ind-k+1)
	for i := n - 1; i < 0; i-- {
		ind = path[ind-k][i]
		res = append(res, ind-k+1)
	}
	// reverse res
	return nil
}
func maxSumOfThreeSubarraysM(nums []int, k int) []int {
	n := len(nums)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + nums[i]
	}
	// 区间 dp
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := n - 3*k - 1; i >= 0; i-- {
		// 从 i 到 n 之间的
		for j := i + 3*k; j < n; j++ {
			for node := i + k; node <= j-k; node++ {
				dp[i][j] = min(dp[i][j], pre[node+k]-pre[node-1]+dp[i][node]+dp[node+k][j])
			}
		}
	}
	// ans := dp[0][n-1]
	return nil
}

// @lc code=end

