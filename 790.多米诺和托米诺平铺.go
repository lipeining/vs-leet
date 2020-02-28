/*
 * @lc app=leetcode.cn id=790 lang=golang
 *
 * [790] 多米诺和托米诺平铺
 */

// @lc code=start
func numTilings(N int) int {
	mod := int(1e9 + 7)
	dp := make([]int, 4)
	dp[0] = 1
	for i := 0; i < N; i++ {
		ndp := make([]int, 4)
		ndp[0] = (dp[0] + dp[3]) % mod
		ndp[1] = (dp[0] + dp[2]) % mod
		ndp[2] = (dp[0] + dp[1]) % mod
		ndp[3] = (dp[0] + dp[1] + dp[2]) % mod
		dp = ndp
	}
	return dp[0]
}

// @lc code=end

// public int numTilings(int N) {
// 	int MOD = 1_000_000_007;
// 	long[] dp = new long[]{1, 0, 0, 0};
// 	for (int i = 0; i < N; ++i) {
// 		long[] ndp = new long[4];
// 		ndp[0b00] = (dp[0b00] + dp[0b11]) % MOD;
// 		ndp[0b01] = (dp[0b00] + dp[0b10]) % MOD;
// 		ndp[0b10] = (dp[0b00] + dp[0b01]) % MOD;
// 		ndp[0b11] = (dp[0b00] + dp[0b01] + dp[0b10]) % MOD;
// 		dp = ndp;
// 	}
// 	return (int) dp[0];
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/domino-and-tromino-tiling/solution/duo-mi-nuo-yu-tuo-mi-nuo-ping-pu-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。