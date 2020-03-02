import "math"

/*
 * @lc app=leetcode.cn id=1130 lang=golang
 *
 * [1130] 叶值的最小代价生成树
 */

// @lc code=start
func mctFromLeafValues(arr []int) int {
	length := len(arr)
	dp := make([][]int, length)
	// 从 0-length 最小值为多少
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		for j := 0; j < length; j++ {
			dp[i][j] = math.MaxInt32
		}
		dp[i][i] = 0
	}
	maxs := make([][]int, length)
	// 表示 区间 i,j 之间最大值
	for i := 0; i < length; i++ {
		maxs[i] = make([]int, length)
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			want := 0
			for k := i; k < j+1; k++ {
				want = max(want, arr[k])
			}
			maxs[i][j] = want
		}
	}
	for l := 1; l < length; l++ {
		for i := 0; i < length-l; i++ {
			j := i + l
			for k := 0; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]+maxs[i][k]*maxs[k+1][j])
			}
		}
	}
	return dp[0][length-1]
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
