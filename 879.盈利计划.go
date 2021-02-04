/*
 * @lc app=leetcode.cn id=879 lang=golang
 *
 * [879] 盈利计划
 */
package main

// func main() {
// 	profitableSchemes(5, 3, []int{2, 2}, []int{2, 3})
// }

// @lc code=start
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	N := len(group)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, minProfit+1)
		dp[i][0] = 1
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	toMod := int(1e9 + 7)
	for i := 1; i <= N; i++ {
		g := group[i-1]
		p := profit[i-1]
		for j := n; j >= g; j-- {
			for k := minProfit; k >= 0; k-- {
				dp[j][k] += dp[j-g][max(k-p, 0)]
				dp[j][k] %= toMod
			}
		}
	}
	return dp[n][minProfit]
	// dp := make([]map[int]int, n+1)
	// for i := 0; i <= n; i++ {
	// 	dp[i] = make(map[int]int)
	// 	dp[i][0] = 1
	// }
	// 	三维动态规划
	// 1，dp[i][j][k]代表考虑前i个profit的情况下，在可用人数为j个情况下，能获取利润至少为k的情况数
	// 2，状态转移方程为：
	// 令g = group[i - 1];
	// 令p = profit[i - 1];
	// 则dp[i][j][k] = dp[i - 1][j][k] + dp[i - 1][j - g][max(k - p, 0)];

	// 作者：da-li-wang
	// 链接：https://leetcode-cn.com/problems/profitable-schemes/solution/c-dong-tai-gui-hua-by-da-li-wang-39/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	// dp[0][0] = 1
	// i 个人 利润 的数量
	// for i := 1; i <= n; i++ {
	// 	for j := 0; j < len(profit); j++ {
	// 		g := group[j]
	// 		pro := profit[j]
	// 		if i >= g {
	// 			for k, v := range dp[i-g] {
	// 				dp[i][pro+k] = v
	// 			}
	// 			fmt.Println(i, g)
	// 		}
	// 	}
	// 	fmt.Println(i, dp[i])
	// }
	// ans := 0
	// for k, v := range dp[n] {
	// 	if k >= minProfit {
	// 		ans += v
	// 	}
	// }
	// fmt.Println(dp[n])
	// return ans
}

// @lc code=end
