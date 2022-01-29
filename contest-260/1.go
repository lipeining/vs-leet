package main

import "math"

func main() {

}
func maximumDifference(nums []int) int {
	n := len(nums)
	ans := -1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				ans = max(ans, nums[j]-nums[i])

			}
		}
	}
	return ans
}
func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	pre := make([][]int64, 2)
	for i := 0; i < 2; i++ {
		pre[i] = make([]int64, n+1)
	}
	for i := 1; i <= n; i++ {
		pre[0][i] = pre[0][i-1] + int64(grid[0][i])
		pre[1][i] = pre[1][i-1] + int64(grid[1][i])
	}
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	ans := int64(math.MaxInt64)
	for i := 1; i <= n; i++ {
		ans = min(ans, max(pre[0][n]-pre[0][i], pre[1][i-1]))
	}
	return ans
}
