/*
 * @lc app=leetcode.cn id=514 lang=golang
 *
 * [514] 自由之路
 */
package main

import (
	"math"
)

// func main() {
// 	fmt.Println("ans", findRotateSteps("godding", "gd"))
// 	// 	"caotmcaataijjxi"
// 	// "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx"
// 	fmt.Println("ans", findRotateSteps("caotmcaataijjxi", "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx"))
// }

// @lc code=start
func findRotateSteps(ring string, key string) int {
	n := len(ring)
	m := len(key)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, m)
		for j := 0; j < m; j++ {
			memo[i][j] = -1
		}
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j == m {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		cost := math.MaxInt32
		for p := 0; p < n; p++ {
			if ring[p] == key[j] {
				cost = min(cost, min(abs(p-i), n-abs(p-i))+dfs(p, j+1))
			}
		}
		// fmt.Println("i, j", i, j)
		memo[i][j] = cost
		return cost
	}
	return m + dfs(0, 0)
}

// @lc code=end
func findRotateStepsDP(ring string, key string) int {
	n := len(ring)
	min := func(a ...int) int {
		res := a[0]
		for i := 1; i < len(a); i++ {
			if a[i] < res {
				res = a[i]
			}
		}
		return res
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	inf := math.MaxInt64 / 2
	n, m := len(ring), len(key)
	pos := [26][]int{}
	for i, c := range ring {
		pos[c-'a'] = append(pos[c-'a'], i)
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = inf
	}
	for _, p := range pos[key[0]-'a'] {
		dp[p] = min(p, n-p) + 1
	}
	for i := 1; i < m; i++ {
		for _, j := range pos[key[i]-'a'] {
			for _, k := range pos[key[i-1]-'a'] {
				// 这里的 k 要保证不覆盖重复
				dp[j] = min(dp[j], dp[k]+min(abs(j-k), n-abs(j-k))+1)
			}
		}
	}
	return dp[n-1]
	// // dp[i][j] 表示 从前往后拼写出 key 的第 i 个字符， ring 的第 j 个字符与 12:00
	// // 方向对齐的最少步数，下标从 0 开始。
	// dp := make([][]int, m)
	// for i := 0; i < m; i++ {
	// 	dp[i] = make([]int, n)
	// 	for j := 0; j < n; j++ {
	// 		dp[i][j] = inf
	// 	}
	// }
	// for _, p := range pos[key[0]-'a'] {
	// 	dp[0][p] = min(p, n-p) + 1
	// }
	// for i := 1; i < m; i++ {
	// 	for _, j := range pos[key[i]-'a'] {
	// 		for _, k := range pos[key[i-1]-'a'] {
	// 			dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), n-abs(j-k))+1)
	// 		}
	// 	}
	// }
	// return min(dp[m-1]...)
	// n := len(ring)
	// var dfs func(r, k string) int
	// dfs = func(r, k string) int {
	// 	if k == "" {
	// 		return 0
	// 	}
	// 	// 第一个字符需要的花费成本，
	// 	// 返回最小的
	// 	char := k[0]
	// 	cost := math.MaxInt32
	// 	for i := 0; i < n; i++ {
	// 		if r[i] == char {
	// 			dis := min(i, n-i)
	// 			cost = min(cost, dis+1+dfs(r[i:]+r[:i], k[1:]))
	// 		}
	// 	}
	// 	return cost
	// }
	// return dfs(ring, key)
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
