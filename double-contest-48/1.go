package main

import (
	"fmt"
	"sort"
)

func main() {
	getMaximumConsecutive([]int{1, 3})
	getMaximumConsecutive([]int{1, 1, 1, 4})
}
func maxScore(nums []int) int {
	m := len(nums)
	// 压缩 dp
	// max := func(a, b int) int {
	// 	if a > b {
	// 		return a
	// 	}
	// 	return b
	// }
	gcds := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		gcds[i] = make([]int, m+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= m; j++ {
			if i != j {
				l := nums[i]
				r := nums[j]
				res := gcd(l, r)
				gcds[i][j] = res
				gcds[j][i] = res
			}
		}
	}
	length := 1 << m
	dp := make([]int, 1<<m)
	for s := 0; s < length; s++ {
		for i := 1; i <= m; i++ {
			for j := 1; j <= m; j++ {
				if i == j {
					continue
				}
				// 需要有这个状态，才可以转换
				if (s&(1<<i) == 0) && (s&(1<<j) == 0) {
					// ns := (s | (1 << i)) | (1 << j)
					// dp[ns] = dp[s] +
				}
			}
		}
	}
	return dp[1<<length-1]
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	sum := 0
	for _, num := range coins {
		sum += num
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, sum+1)
	}
	ans := 1
	for i := 1; i <= n; i++ {
		num := coins[i-1]
		dp[i][num] = 1
	}
	for i := 1; i <= n; i++ {
		num := coins[i-1]
		for j := 0; j <= sum; j++ {

		}
	}
	fmt.Println(dp)
	fmt.Println("anssss", ans)
	return ans
}
func getMaximumConsecutive2(coins []int) int {
	sort.Ints(coins)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(coins)
	dp := make([]map[int]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make(map[int]int)
		dp[0][0] = 1
	}
	ans := 1
	for i := 1; i <= n; i++ {
		num := coins[i-1]
		dp[i][num] = 1
	}
	for i := 1; i <= n; i++ {
		num := coins[i-1]
		for j := 0; j < i; j++ {
			for k, v := range dp[j] {
				next := k + num
				if dp[j][next-1] != 0 {
					dp[i][next] = max(dp[i][next], v+1)
					ans = max(ans, dp[i][next])
				}
			}
		}
		fmt.Println("i", i, dp[i])
	}
	fmt.Println(dp)
	fmt.Println("anssss", ans)
	return ans
}
