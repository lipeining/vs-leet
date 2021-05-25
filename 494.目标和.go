/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < S || (S+sum)%2 != 0 {
		return 0
	}
	target := (S + sum) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	// 前 i 个数，0,1 的取法，可以得到 target 的方法数
	// target 从大到小
	for i := 1; i <= n; i++ {
		num := nums[i-1]
		for j := target; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[target]
	// n := len(nums)
	// ans := 0
	// // dp 转为 sum/2 的零一背包
	// var dfs func(index int, sum int)
	// dfs = func(index int, sum int) {
	// 	if index == n {
	// 		if sum == S {
	// 			ans++
	// 		}
	// 		return
	// 	}
	// 	dfs(index+1, sum+nums[index])
	// 	dfs(index+1, sum-nums[index])
	// }
	// dfs(0, 0)
	// return ans
}
func findTargetSumWaysDP(nums []int, S int) int {
	if S > 1000 {
		return 0
	}
	sum := 2001
	length := len(nums)
	// dp[i][j] 第 i 次，和为 j 的方案数
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, sum)
	}
	// dp[i][j + nums[i] + 1000] += dp[i - 1][j + 1000]
	// dp[i][j - nums[i] + 1000] += dp[i - 1][j + 1000]

	// 作者：LeetCode
	// 链接：https://leetcode-cn.com/problems/target-sum/solution/mu-biao-he-by-leetcode/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	dp[0][nums[0]+1000] = 1
	dp[0][-nums[0]+1000] += 1
	for i := 1; i < length; i++ {
		for sum := -1000; sum <= 1000; sum++ {
			if dp[i-1][sum+1000] > 0 {
				dp[i][sum+nums[i]+1000] += dp[i-1][sum+1000]
				dp[i][sum-nums[i]+1000] += dp[i-1][sum+1000]
			}
		}
	}
	return dp[length-1][S+1000]
}

// @lc code=end
