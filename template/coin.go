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

// 因为硬币可以重复使用，因此这是一个完全背包问题。完全背包只需要将 0-1 背包的逆序遍历 dp 数组改为正序遍历即可。

// public int coinChange(int[] coins, int amount) {
//     int[] dp = new int[amount + 1];
//     for (int coin : coins) {
//         for (int i = coin; i <= amount; i++) { //将逆序遍历改为正序遍历
//             if (i == coin) {
//                 dp[i] = 1;
//             } else if (dp[i] == 0 && dp[i - coin] != 0) {
//                 dp[i] = dp[i - coin] + 1;

//             } else if (dp[i - coin] != 0) {
//                 dp[i] = Math.min(dp[i], dp[i - coin] + 1);
//             }
//         }
//     }
//     return dp[amount] == 0 ? -1 : dp[amount];
// }

// 求解顺序的完全背包问题时，对物品的迭代应该放在最里层，
// s = "leetcode",
// dict = ["leet", "code"].
// Return true because "leetcode" can be segmented as "leet code".
// 对背包的迭代放在外层，只有这样才能让物品按一定顺序放入背包中。

// public boolean wordBreak(String s, List<String> wordDict) {
//     int n = s.length();
//     boolean[] dp = new boolean[n + 1];
//     dp[0] = true;
//     for (int i = 1; i <= n; i++) {
//         for (String word : wordDict) {
// 对物品的迭代应该放在最里层
//             int len = word.length();
//             if (len <= i && word.equals(s.substring(i - len, i))) {
//                 dp[i] = dp[i] || dp[i - len];
//             }
//         }
//     }
//     return dp[n];
// }
