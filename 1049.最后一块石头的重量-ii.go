/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
func lastStoneWeightII(stones []int) int {
	length := len(stones)
	sum := 0
	for i := 0; i < length; i++ {
		sum += stones[i]
	}
	half := sum / 2
	// 抄答案的
	/* 定义 dp[i] 重量上限为 i 时背包所能装载的最大石头重量 */
	dp := make([]int, half+1)
	for i := 0; i < length; i++ {
		stone := stones[i]
		for j := half; j >= stone; j-- {
			dp[j] = max(dp[j], dp[j-stone]+stone)
		}
	}
	return sum - 2*dp[half]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
