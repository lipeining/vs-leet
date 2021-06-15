package main

import (
	"fmt"
	"sort"
)

func main() {

}
func largestMagicSquare(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 二分加暴力
	l := 2
	r := min(m, n)
	// 行的前缀和
	row := make([][]int, m)
	for i := 0; i < m; i++ {
		row[i] = make([]int, n+1)
		for j := 0; j < n; j++ {
			row[i][j+1] = row[i][j] + grid[i][j]
		}
	}
	// 列的前缀和
	col := make([][]int, n)
	for i := 0; i < n; i++ {
		col[i] = make([]int, m+1)
		for j := 0; j < m; j++ {
			col[i][j+1] = col[i][j] + grid[j][i]
		}
	}
	pass := func(size int) bool {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if i+size > m || j+size > n {
					continue
				}
				// 计算一个方块
				want := row[i][j+size] - row[i][j]
				flag := true
				d1, d2 := 0, 0
				for k := 0; k < size; k++ {
					r := row[i+k][j+size] - row[i+k][j]
					c := col[j+k][i+size] - col[j+k][i]
					d1 += grid[i+k][j+k]
					d2 += grid[i+k][j+size-1-k]
					if r != want || c != want {
						flag = false
						break
					}
				}
				fmt.Println(d1, d2, flag)
				if d1 != want || d2 != want {
					continue
				}
				if flag {
					return true
				}
			}
		}
		return false
	}
	for l < r {
		mid := (l + r + 1) >> 1
		if pass(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if pass(l) {
		return l
	}
	return l - 1
}
func chalkReplacer(chalk []int, k int) int {
	n := len(chalk)
	sum := 0
	for i := 0; i < n; i++ {
		sum += chalk[i]
	}
	k = k % sum
	i := 0
	for k >= chalk[i] {
		k -= chalk[i]
		i++
		i %= n
	}
	return i
}
func isCovered(ranges [][]int, left int, right int) bool {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] != ranges[j][0] {
			return ranges[i][0] < ranges[j][0]
		}
		return ranges[i][1] < ranges[j][1]
	})
	i := 0
	n := len(ranges)
	for x := left; x <= right; x++ {
		flag := false
		for i < n {
			if ranges[i][0] <= x && ranges[i][1] >= x {
				flag = true
				break
			} else {
				i++
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
