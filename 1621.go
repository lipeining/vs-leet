package main

import "fmt"

// func main() {
	// solve(4, 2) //5
	// solve(3, 1)  //3
	// solve(30, 7) // 796297179
	// solve(5, 3)  //7
	// solve(3, 2)  //1
// }

func solve(n int, k int) int {
	// mod := int(1e9 + 7)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
	}
	// 到下标为 i 的点，能否划分 j 段
	// 前 x 个为 j-1 段
	// min := func(a, b int) int {
	// 	if a < b {
	// 		return a
	// 	}
	// 	return b
	// }
	ans := 0
	// fmt.Println(dp)
	fmt.Println("anssss", ans)
	return ans
}
