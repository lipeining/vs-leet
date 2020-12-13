/*
 * @lc app=leetcode.cn id=486 lang=golang
 *
 * [486] 预测赢家
 */
package main

import (
	"strconv"
)

// func main() {
// 	fmt.Println("ans", PredictTheWinner([]int{1, 1, 1}))
// 	fmt.Println("ans", PredictTheWinner([]int{1, 567, 1, 1, 99, 100}))
// 	fmt.Println("ans", PredictTheWinner([]int{1, 5, 2}))
// 	fmt.Println("ans", PredictTheWinner([]int{1, 5, 233, 7}))
// }

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func PredictTheWinner(nums []int) bool {
	// var dfs func(start, end int, turn int) int
	// dfs = func(start, end int, turn int) int {
	// 	if start == end {
	// 		return nums[start] * turn
	// 	}
	// 	scoreStart := nums[start]*turn + dfs(start+1, end, -turn)
	// 	scoreEnd := nums[end]*turn + dfs(start, end-1, -turn)
	// 	return max(scoreStart*turn, scoreEnd*turn) * turn
	// }
	// return dfs(0, len(nums)-1, 1) >= 0
	n := len(nums)
	// 当前玩家能拿到的最大分数之差
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}
	// dp[i][j] = max(nums[i]- dp[i+1][j], nums[j]- dp[i][j-1])
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}
func PredictTheWinner2(nums []int) bool {
	// memory dfs
	n := len(nums)
	// 在只有 i,j 之间的数组时，第一个是否能赢?
	memo := make(map[string]bool)
	var dfs func(i, j int, left, right int) bool
	dfs = func(i, j int, left, right int) bool {
		key := strconv.Itoa(i) + "-" + strconv.Itoa(j)
		// fmt.Println("i, j, memo", i, j, memo)
		if res, ok := memo[key]; ok {
			return res
		}
		if i == j {
			return left+nums[i] >= right
		}
		if i == j-1 {
			if left+nums[i] >= right+nums[j] {
				return true
			}
			if left+nums[j] >= right+nums[i] {
				return true
			}
			// fmt.Println("i, j, left, right", i, j, left, right)
			return false
		}
		// 不止两个元素，那么需要分类讨论
		if !dfs(i+1, j, left-nums[i], right) {
			// 选左边且之后不是输家
			memo[key] = true
			return true
		}
		if !dfs(i, j-1, left-nums[j], right) {
			// 选左边且之后不是输家
			memo[key] = true
			return true
		}
		memo[key] = false
		// fmt.Println("chooese i, j, left, right", i, j, left, right)
		return false
	}
	return dfs(0, n-1, 0, 0)
}

// @lc code=end
