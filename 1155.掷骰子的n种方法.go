/*
 * @lc app=leetcode.cn id=1155 lang=golang
 *
 * [1155] 掷骰子的N种方法
 */

// @lc code=start
func numRollsToTarget(d int, f int, target int) int {
	var dp [31][1001]int
	toMod := int(1e9 + 7)
	for i := 1; i <= f && i <= target; i++ {
		dp[1][i] = 1
	}
	// targetMax := d * f
	for i := 2; i <= d; i++ {
		for j := i; j <= target; j++ {
			for k := 1; k <= f && k <= j; k++ {
				dp[i][j] = (dp[i][j] + dp[i-1][j-k]) % toMod
			}
		}
	}
	return dp[d][target]
}

// @lc code=end
// public int numRollsToTarget(int d, int f, int target) {
//     int[][] dp = new int[31][1001];
//     int min = Math.min(f, target);
//     for (int i = 1; i <= min; i++) {
//       dp[1][i] = 1;
//     }
//     int targetMax = d * f;
//     for (int i = 2; i <= d; i++) {
//       for (int j = i; j <= targetMax; j++) {
//         for (int k = 1; j - k >= 0 && k <= f; k++) {
//           dp[i][j] = (dp[i][j] + dp[i - 1][j - k]) % MOD;
//         }
//       }
//     }
//     return dp[d][target];
//   }

// 作者：maverickbytes
// 链接：https://leetcode-cn.com/problems/number-of-dice-rolls-with-target-sum/solution/zuo-ti-guo-cheng-ji-lu-dpjie-fa-by-maverickbytes/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
