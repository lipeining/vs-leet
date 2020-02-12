import "math"

/*
 * @lc app=leetcode.cn id=375 lang=golang
 *
 * [375] 猜数字大小 II
 */

// @lc code=start
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for len := 2; len <= n; len++ {
		for start := 1; start <= n-len+1; start++ {
			minres := math.MaxInt32
			for piv := start; piv < start+len-1; piv++ {
				res := piv + max(dp[start][piv-1], dp[piv+1][start+len-1])
				minres = min(res, minres)
			}
			// 应该先写下面这句，再推算出 start <= n-len+1
			dp[start][start+len-1] = minres
		}
	}
	return dp[1][n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

// public class Solution {
//     public int getMoneyAmount(int n) {
//         int[][] dp = new int[n + 1][n + 1];
//         for (int len = 2; len <= n; len++) {
//             for (int start = 1; start <= n - len + 1; start++) {
//                 int minres = Integer.MAX_VALUE;
//                 for (int piv = start; piv < start + len - 1; piv++) {
//                     int res = piv + Math.max(dp[start][piv - 1], dp[piv + 1][start + len - 1]);
//                     minres = Math.min(res, minres);
//                 }
//                 dp[start][start + len - 1] = minres;
//             }
//         }
//         return dp[1][n];
//     }
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/solution/cai-shu-zi-da-xiao-ii-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。