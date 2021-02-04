/*
 * @lc app=leetcode.cn id=877 lang=golang
 *
 * [877] 石子游戏
 */

// @lc code=start
func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = piles[i]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[j] = max(piles[i]-dp[j], piles[j]-dp[j-1])
		}
	}
	return dp[n-1] > 0
	// N := len(piles)
	// dp := make([][]int, N+2)
	// for i := 0; i < len(dp); i++ {
	// 	dp[i] = make([]int, N+2)
	// }
	// for size := 1; size <= N; size++ {
	// 	for i := 0; i+size <= N; i++ {
	// 		j := i + size - 1
	// 		parity := (j + i + N) % 2
	// 		if parity == 1 {
	// 			max := piles[i] + dp[i+2][j+1]
	// 			if max < piles[j]+dp[i+1][j] {
	// 				max = piles[j] + dp[i+1][j]
	// 			}
	// 			dp[i+1][j+1] = max
	// 		} else {
	// 			min := -piles[i] + dp[i+2][j+1]
	// 			if min > (-piles[j] + dp[i+1][j]) {
	// 				min = (-piles[j] + dp[i+1][j])
	// 			}
	// 			dp[i+1][j+1] = min

	// 		}
	// 	}
	// }
	// return dp[1][N] > 0
}

// @lc code=end

// public boolean stoneGame(int[] piles) {
// 	int N = piles.length;

// 	// dp[i+1][j+1] = the value of the game [piles[i], ..., piles[j]].
// 	int[][] dp = new int[N+2][N+2];
// 	for (int size = 1; size <= N; ++size)
// 		for (int i = 0; i + size <= N; ++i) {
// 			int j = i + size - 1;
// 			int parity = (j + i + N) % 2;  // j - i - N; but +x = -x (mod 2)
// 			if (parity == 1)
// 				dp[i+1][j+1] = Math.max(piles[i] + dp[i+2][j+1], piles[j] + dp[i+1][j]);
// 			else
// 				dp[i+1][j+1] = Math.min(-piles[i] + dp[i+2][j+1], -piles[j] + dp[i+1][j]);
// 		}

// 	return dp[1][N] > 0;
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/stone-game/solution/shi-zi-you-xi-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。