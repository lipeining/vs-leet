import "fmt"

/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		coin := coins[i]
		for j := 1; j <= amount; j++ {
			if j >= coin {
				dp[j] += dp[j-coin]
			}
		}
	}
	fmt.Println(dp)
	return dp[amount]

	// // 自底向上
	// dp[0] = 1
	// for j := 1; j <= amount; j++ {
	// 	sum := 0
	// 	for i := 0; i < len(coins); i++ {
	// 		coin := coins[i]
	// 		left := j - coin
	// 		if left >= 0 {
	// 			sum += dp[j-coin]
	// 		}
	// 	}
	// 	dp[j] = sum
	// }

}

// @lc code=end
