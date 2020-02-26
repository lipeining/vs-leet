/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] 地下城游戏
 */

// @lc code=start
func calculateMinimumHP(dungeon [][]int) int {
	// 自底向上的算法，从右下角开始
	// 自顶向下的算法，从左上角开始
	m := len(dungeon)
	if m == 0 {
		return 1
	}
	n := len(dungeon[0])
	if n == 0 {
		return 1
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = 1 - min(dungeon[m-1][n-1], 0)
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = max(dp[i+1][n-1]-dungeon[i][n-1], 1)
	}
	for i := n - 2; i >= 0; i-- {
		dp[m-1][i] = max(dp[m-1][i+1]-dungeon[m-1][i], 1)
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = max(min(dp[i+1][j], dp[i][j+1])-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
	// dp := make([][]int, m+1)
	// for i := 0; i <= m; i++ {
	// 	dp[i] = make([]int, n+1)
	// }
	// for i := m; i >= 1; i-- {
	// 	for j := n; j >= 1; j-- {

	// 	}
	// }
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//   /* dp[i][j]:到达第i行第j个格子前需要的最低hp */
//     /* dp[i][j] + dungeon[i][j] >= dp[i + 1][j] */
//     /* dp[i][j] + dungeon[i][j] >= dp[i][j + 1] */
//     /* 上述两个式子满足其一即可，以及要保证任何时候的dp都是大于0的 */

//     dp[R - 1][C - 1] = 1 - min(dungeon[R - 1][C - 1], 0);

//     for (int i = R - 2; i >= 0; i--)           /* 初始化最后一列 */
//         dp[i][C - 1] = max(dp[i + 1][C - 1] - dungeon[i][C - 1], 1);

//     for (int i = C - 2; i >= 0; i--)           /* 初始化最后一行 */
//         dp[R - 1][i] = max(dp[R - 1][i + 1] - dungeon[R - 1][i], 1);

//     for (int i = R - 2; i >= 0; i--)
//         for (int j = C - 2; j >= 0; j--)
//             dp[i][j] = max(min(dp[i + 1][j], dp[i][j + 1]) - dungeon[i][j], 1);

//     return dp[0][0];

// 作者：NPU_V
// 链接：https://leetcode-cn.com/problems/dungeon-game/solution/dong-tai-gui-hua-cyu-yan-by-npu_v/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// @lc code=end
