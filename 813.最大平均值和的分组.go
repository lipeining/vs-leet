/*
 * @lc app=leetcode.cn id=813 lang=golang
 *
 * [813] 最大平均值和的分组
 */
package main

import "fmt"

// func main() {
// 	largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3)
// }

// @lc code=start
func largestSumOfAverages(A []int, K int) float64 {
	n := len(A)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + A[i]
	}
	// 区间 dp
	// 前x个数为k-1个区间，后面的段为一个区间
	// dp[n][K] 为答案
	max := func(a, b float64) float64 {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]float64, K+1)
		dp[i][1] = float64(pre[i+1]) / float64((i + 1))
	}
	// fmt.Println(dp)
	for x := 2; x <= K; x++ {
		for i := 0; i < n; i++ {
			for j := 0; j < i; j++ {
				// j之前为 k-1 个区间， j:i 为一个区间
				left := dp[j][x-1]
				right := float64(pre[i+1]-pre[j+1]) / float64(i-j)
				// fmt.Println("i,j,x", i, j, x, "left", left, "right", right)
				dp[i][x] = max(dp[i][x], left+right)
			}
		}
		// fmt.Println("区间数", x, dp)
	}
	ans := dp[n-1][K]
	fmt.Println("ansssss", ans)
	return ans
}

// @lc code=end
