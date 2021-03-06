/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	// length := n
	// dp := make([][]int, length)
	// // 从 0-length 最小值为多少
	// for i := 0; i < length; i++ {
	// 	dp[i] = make([]int, length)
	// 	// for j := 0; j < length; j++ {
	// 	// 	dp[i][j] = 0
	// 	// }
	// 	dp[i][i] = 1
	// }
	// for l := 1; l < length; l++ {
	// 	for i := 0; i < length-l; i++ {
	// 		j := i + l
	// 		sum := 0
	// 		for k := 0; k < j; k++ {
	// 			sum += dp[i][k] * dp[k+1][j]
	// 		}
	// 		dp[i][j] = sum
	// 	}
	// }
	// return dp[0][length-1]

	G := make([]int, n+1)
	G[0] = 1
	G[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
	//     int[] G = new int[n + 1];
	//     G[0] = 1;
	//     G[1] = 1;

	//     for (int i = 2; i <= n; ++i) {
	//       for (int j = 1; j <= i; ++j) {
	//         G[i] += G[j - 1] * G[i - j];
	//       }
	//     }
	//     return G[n];

	// 作者：LeetCode
	// 链接：https://leetcode-cn.com/problems/unique-binary-search-trees/solution/bu-tong-de-er-cha-sou-suo-shu-by-leetcode/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}

// 超时了
// func numTrees(n int) int {
//     if n == 0 {
// 		return 0
// 	}
// 	return helper(1, n)
// }
// func helper(start int, end int) int {
// 	if start > end {
// 		return 1
// 	}
// 	ans := 0
// 	for i:=start;i<=end;i++ {
// 		lefts := helper(start, i-1)
// 		rights := helper(i+1, end)
// 		multi := lefts * rights
// 		ans += multi
// 	}
// 	return ans
// }
// @lc code=end
