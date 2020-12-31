/*
 * @lc app=leetcode.cn id=740 lang=golang
 *
 * [740] 删除与获得点数
 */

// @lc code=start
func deleteAndEarn(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		cnt[nums[i]]++
	}
	size := nums[n-1]
	dp := make([]int, size+1)
	dp[1] = cnt[1]
	for i := 2; i <= size; i++ {
		dp[i] = max(dp[i-2]+cnt[i]*i, dp[i-1])
	}
	return dp[size]
}

// @lc code=end

