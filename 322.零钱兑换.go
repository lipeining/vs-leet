/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
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

	// 正确解答：2
	dp := make([]int, amount+1)
	for j := 1; j <= amount; j++ {
		min := 2 * amount
		for i := 0; i < len(coins); i++ {
			coin := coins[i]
			if coin > j {
				continue
			}
			if dp[j-coin] == -1 {
				continue
			}
			if min > dp[j-coin]+1 {
				min = dp[j-coin] + 1
			}
		}
		if min == 2*amount {
			dp[j] = -1
		} else {
			dp[j] = min
		}
	}
	return dp[amount]

	// 正确解答：1
	// dp := make([]int, amount+1)
	// for j := 0; j <= amount; j++ {
	// 	dp[j] = math.MaxInt32
	// }
	// dp[0] = 0

	// for i := 0; i < len(coins); i++ {
	// 	coin := coins[i]
	// 	for j := coin; j <= amount; j++ {
	// 		min := dp[j-coin] + 1
	// 		if min > dp[j] {
	// 			min = dp[j]
	// 		}
	// 		dp[j] = min
	// 	}
	// }
	// if dp[amount] == math.MaxInt32 {
	// 	return -1
	// }
	// return dp[amount]

	// 错误解答：1
	// for i := 0; i < len(coins); i++ {
	// 	for j := 1; j*coins[i] <= amount; j++ {
	// 		dp[coins[i]*j] = j
	// 	}
	// }
	// fmt.Println(dp)
	// for i := 0; i < len(coins); i++ {
	// 	coin := coins[i]
	// 	for j := amount; j >= coin; j-- {
	// 		if j >= coin {
	// 			min := dp[j-coin] + 1
	// 			if min > dp[j] {
	// 				min = dp[j]
	// 			}
	// 			dp[j] = min
	// 		}
	// 	}
	// }
	// fmt.Println(dp)
	// if dp[amount] == math.MaxInt32 {
	// 	return -1
	// }
	// return dp[amount]
}

// @lc code=end
