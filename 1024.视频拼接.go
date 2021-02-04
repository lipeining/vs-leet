/*
 * @lc app=leetcode.cn id=1024 lang=golang
 *
 * [1024] 视频拼接
 */
package main

import (
	"fmt"
	"math"
	"sort"
)

// func main() {
// 	videoStitching([][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 10)
// 	videoStitching([][]int{{0, 1}, {1, 2}}, 5)
// 	videoStitching([][]int{{0, 4}, {2, 8}}, 5)
// }

// @lc code=start
func videoStitching(clips [][]int, T int) int {
	n := len(clips)
	dp := make([]int, T+1)
	for i := 0; i <= T; i++ {
		dp[i] = math.MaxInt32
	}
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0]
	})
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 背包 dp
	dp[0] = 0
	for i := 0; i < n; i++ {
		clip := clips[i]
		for j := clip[0]; j <= clip[1] && j <= T; j++ {
			dp[j] = min(dp[j], dp[clip[0]]+1)
		}
	}
	ans := dp[T]
	fmt.Println("anssssss", ans)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// @lc code=end
