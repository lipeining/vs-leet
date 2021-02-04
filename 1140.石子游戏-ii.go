/*
 * @lc app=leetcode.cn id=1140 lang=golang
 *
 * [1140] 石子游戏 II
 */
package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	stoneGameII([]int{2, 7, 9, 4, 4})
// }

// @lc code=start
func stoneGameII(piles []int) int {
	n := len(piles)
	// min := func(a, b int) int {
	// 	if a < b {
	// 		return a
	// 	}
	// 	return b
	// }
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 当前用户比另外一个可以多拿的数量
	memo := make(map[string]int)
	var dfs func(i int, m int) int
	dfs = func(i int, m int) int {
		if i == n {
			return 0
		}
		key := strconv.Itoa(i) + "-" + strconv.Itoa(m)
		if memo[key] != 0 {
			return memo[key]
		}
		next := -1000 * 10000
		take := 0
		nextM := m
		for x := 1; x <= 2*m && i+x <= n; x++ {
			take += piles[i+x-1]
			nextM = max(nextM, x)
			next = max(next, take-dfs(i+x, nextM))
		}
		memo[key] = next
		// fmt.Println("current index ", i, m, "has next ", next)
		return next
	}
	more := dfs(0, 1)
	total := 0
	for _, num := range piles {
		total += num
	}
	ans := (total + more) / 2
	fmt.Println("ansssss", ans)
	return ans
}

// @lc code=end
