/*
 * @lc app=leetcode.cn id=754 lang=golang
 *
 * [754] 到达终点数字
 */
package main

import (
	"fmt"
	"math"
)

// func main() {
// 	reachNumber(3)
// 	reachNumber(2)
// }

// @lc code=start
func reachNumber(target int) int {
	memo := make(map[int]int)
	ans := math.MaxInt32
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var dfs func(index, move int, path int)
	dfs = func(index, move int, path int) {
		// 在 index 的位置，可以向左或者右
		if oldp, ok := memo[index]; ok {
			if oldp <= path {
				return
			}
		}
		memo[index] = path
		if index == target {
			ans = min(ans, path)
			return
		}
		if index > target {
			dfs(index-move, move+1, path+1)
		} else {
			if target-index < move {
				dfs(index-move, move+1, path+1)
			}
			dfs(index+move, move+1, path+1)
		}
	}
	dfs(0, 1, 0)
	// fmt.Println("memo", memo)
	fmt.Println("ansssssssss", ans)
	return ans
}

// @lc code=end
