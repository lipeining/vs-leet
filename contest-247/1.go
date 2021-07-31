package main

import (
	"fmt"
	"math"
)
func wonderfulSubstrings(word string) int64 {
	n := len(word)
	cnt := make([]int, 10)
	l := 0
	r := 0
	now := 0// 当前奇数元素个数
	// 要么枚举 10 个元素，当他们为奇数个时的答案。
	for r < n {
		cnt[word[r]-'a']++

	}
}
func main() {
	rotateGrid([][]int{{40, 10}, {30, 20}}, 1)
	rotateGrid([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, 2)
	rotateGrid([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}}, 2)
	rotateGrid([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}, 2)
	// 1 2 3 4         5
	// 14 15 16 17       6
	// 13   20    19  18      7
	// 12 11 10 9 8
	
}
func rotateGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		fmt.Println(grid[i])
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	level := min(m, n) / 2
	fmt.Println("level", level)
	for l := 0; l < level; l++ {
		// 每一个 level 的长度为
		curm := m - l*2
		curn := n - l*2
		size := (curm+curn)*2 - 4
		// k 与 size 的关系
		k %= size
		// 逆时针 k 次
		list := make([]int, 0)
		r0, c0 := l, l
		edgem := r0 + curm
		edgen := c0 + curn
		for j := c0; j < edgen-1; j++ {
			list = append(list, grid[r0][j])
		}
		for i := r0; i < edgem-1; i++ {
			list = append(list, grid[i][edgen-1])
		}
		for j := edgen - 1; j > c0; j-- {
			list = append(list, grid[edgem-1][j])
		}
		for i := edgem - 1; i > r0; i-- {
			list = append(list, grid[i][c0])
		}
		fmt.Println(r0, c0, edgem, edgen, list)
		step := k
		for j := c0; j < edgen-1; j++ {
			grid[r0][j] = list[step]
			step++
			step %= len(list)
		}
		for i := r0; i < edgem-1; i++ {
			grid[i][edgen-1] = list[step]
			step++
			step %= len(list)
		}
		for j := edgen - 1; j > c0; j-- {
			grid[edgem-1][j] = list[step]
			step++
			step %= len(list)
		}
		for i := edgem - 1; i > r0; i-- {
			grid[i][c0] = list[step]
			step++
			step %= len(list)
		}
	}
	fmt.Println(grid)
	return grid
}
func maxProductDifference(nums []int) int {
	n := len(nums)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	left := 0
	right := math.MaxInt32
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			multi := nums[i] * nums[j]
			left = max(left, multi)
			right = min(right, multi)
		}
	}
	return left - right
}
