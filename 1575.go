package main

import "fmt"

// func main() {
// 	countRoutes([]int{2, 3, 6, 8, 4}, 1, 3, 5)
// 	countRoutes([]int{4, 3, 1}, 1, 0, 6)
// 	countRoutes([]int{5, 2, 1}, 0, 2, 3)
// 	countRoutes([]int{2, 1, 5}, 0, 0, 3)
// 	countRoutes([]int{1, 2, 3}, 0, 2, 40)
// }
func countRoutes(locations []int, start, finish, fuel int) int {
	n := len(locations)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, fuel+1)
	}
	mod := int(1e9 + 7)
	abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}
	dp[start][fuel] = 1
	// 在 f 的数量时，从 i 到 j 的方法
	for f := fuel; f >= 1; f-- {
		for i := 0; i < n; i++ {
			if dp[i][f] != 0 {
				for j := 0; j < n; j++ {
					if i == j {
						continue
					}
					need := abs(locations[i] - locations[j])
					if f >= need {
						dp[j][f-need] += dp[i][f]
						dp[j][f-need] %= mod
					}
				}
			}
		}
	}
	ans := 0
	for i := 0; i <= fuel; i++ {
		ans += dp[finish][i]
		ans %= mod
	}
	// fmt.Println(dp)
	// fmt.Println(ans)
	return ans
}
func solve1575old(locations []int, start, finish, fuel int) int {
	n := len(locations)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, fuel+1)
	}
	mod := int(1e9 + 7)
	abs := func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}
	var dfs func(s, f, cnt int)
	dfs = func(s, f, cnt int) {
		dp[s][f] += cnt
		dp[s][f] %= mod
		for i := 0; i < n; i++ {
			if i == s {
				continue
			}
			need := abs(locations[s] - locations[i])
			if f >= need {
				fmt.Println("from", s, "to", i, "left", f-need)
				dfs(i, f-need, cnt)
			}
		}
	}
	dfs(start, fuel, 1)
	ans := 0
	for i := 0; i <= fuel; i++ {
		ans += dp[finish][i]
		ans %= mod
	}
	fmt.Println(dp)
	fmt.Println(ans)
	return ans
}
