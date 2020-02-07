package template

import (
	"fmt"
	"math"
)

// pack 01 的变形 完全背包，即不限制物品使用次数，选择次数
// 322
func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}

	// 正确解答：1
	dp := make([]int, amount+1)
	for j := 0; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}
	dp[0] = 0

	for i := 0; i < len(coins); i++ {
		coin := coins[i]
		for j := coin; j <= amount; j++ {
			min := dp[j-coin] + 1
			if min > dp[j] {
				min = dp[j]
			}
			dp[j] = min
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// 518
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
}
