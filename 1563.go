package main

import "fmt"

// func main() {
// 	stoneGameV([]int{6, 2, 3, 4, 5, 5})
// 	stoneGameV([]int{7, 7, 7, 7, 7, 7, 7})
// 	stoneGameV([]int{4})
// }
func stoneGameV(stoneValues []int) int {
	n := len(stoneValues)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + stoneValues[i]
	}
	memo := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			memo[i][j] = -1
		}
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	fmt.Println(pre)
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if memo[l][r] != -1 {
			return memo[l][r]
		}
		if l == r {
			memo[l][r] = 0
			return 0
		}
		ret := 0
		for k := l; k < r; k++ {
			left := pre[k] - pre[l-1]
			right := pre[r] - pre[k]
			if left > right {
				ret = max(ret, right+dfs(k+1, r))
			} else if left < right {
				ret = max(ret, left+dfs(l, k))
			} else {
				ret = max(ret, max(left+dfs(l, k), right+dfs(k+1, r)))
			}
		}
		memo[l][r] = ret
		return memo[l][r]
	}
	ans := dfs(1, n)
	fmt.Println("anssss", ans)
	return ans
}
