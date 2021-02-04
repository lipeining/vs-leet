/*
 * @lc app=leetcode.cn id=1043 lang=golang
 *
 * [1043] 分隔数组以得到最大和
 */
package main

// import "fmt"

// func main() {
// 	maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3)
// 	maxSumAfterPartitioning([]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4)
// 	maxSumAfterPartitioning([]int{1}, 1)
// }

// @lc code=start
func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// dp[i] 为到第 i 位，最大值为多少
	for i := 1; i <= n; i++ {
		cur := arr[i-1]
		for x := 1; x <= k && i-x >= 0; x++ {
			cur = max(cur, arr[i-x])
			dp[i] = max(dp[i], dp[i-x]+cur*x)
			// fmt.Println("i", i, "x", x, "cur", cur, dp[i])
		}
	}
	return dp[n]
}

// @lc code=end
