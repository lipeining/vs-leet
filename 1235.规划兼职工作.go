/*
 * @lc app=leetcode.cn id=1235 lang=golang
 *
 * [1235] 规划兼职工作
 */
package main

import "sort"

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60})
// 	jobScheduling([]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4})
// 	jobScheduling([]int{6, 15, 7, 11, 1, 3, 16, 2}, []int{19, 18, 19, 16, 10, 8, 19, 8}, []int{2, 9, 1, 19, 5, 7, 3, 19})
// 	// [6,15,7,11,1,3,16,2]
// 	// [19,18,19,16,10,8,19,8]
// 	// [2,9,1,19,5,7,3,19]

// 	// [4,2,4,8,2]
// 	// [5,5,5,10,8]
// 	// [1,2,8,10,4]
// }

// @lc code=start
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	type obj struct {
		s int
		e int
		p int
	}
	n := len(profit)
	list := make([]obj, n)
	for i := 0; i < n; i++ {
		list[i] = obj{
			s: startTime[i],
			e: endTime[i],
			p: profit[i],
		}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].e < list[j].e
	})
	// fmt.Println(list)
	dp := make([]int, n)
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		left := 0
		if i > 0 {
			left = dp[i-1]
		}
		dp[i] = max(left, list[i].p)
		for j := i - 1; j >= 0; j-- {
			if list[i].s >= list[j].e {
				dp[i] = max(dp[i], dp[j]+list[i].p)
				break
			}
		}
		ans = max(ans, dp[i])
	}
	// fmt.Println(dp, ans)
	return ans
}

// @lc code=end
