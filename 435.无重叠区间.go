/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */
package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})
// 	eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}})
// 	eraseOverlapIntervals([][]int{{1, 2}, {2, 3}})
// }

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] > intervals[j][0] {
			return false
		}
		return intervals[i][1] < intervals[j][1]
	})
	// fmt.Println(intervals)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	length := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if intervals[i][0] >= intervals[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		length = max(length, dp[i])
	}
	cnt := n - length
	fmt.Println("anssssssssssssss", cnt)
	return cnt
	// right := intervals[0][1]
	// cnt := 0
	// for i := 1; i < n; i++ {
	// 	// fmt.Println("on i", i, right)
	// 	if right > intervals[i][0] {
	// 		cnt++
	// 	} else {
	// 		right = intervals[i][1]
	// 	}
	// }
	// fmt.Println("anssssssssssssss", cnt)
	// return cnt
}

// @lc code=end
