/*
 * @lc app=leetcode.cn id=1314 lang=golang
 *
 * [1314] 矩阵区域和
 */
package main

import "fmt"

// func main() {
// 	matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1)
// 	matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2)
// }

// @lc code=start
func matrixBlockSum(mat [][]int, K int) [][]int {
	m := len(mat)
	n := len(mat[0])
	pre := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		pre[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1] + mat[i-1][j-1]
		}
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	indexValue := func(x, y int) int {
		x = max(min(x, m), 0)
		y = max(min(y, n), 0)
		return pre[x][y]
	}
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = indexValue(i+K+1, j+K+1) + indexValue(i-K, j-K) - indexValue(i-K, j+K+1) - indexValue(i+K+1, j-K)
		}
	}
	return ans
}
func matrixBlockSum2(mat [][]int, K int) [][]int {
	m := len(mat)
	n := len(mat[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				mat[i][j] = mat[i][j-1] + mat[i][j]
			} else if j == 0 {
				mat[i][j] = mat[i-1][j] + mat[i][j]
			} else {
				mat[i][j] = mat[i-1][j] + mat[i][j-1] + mat[i][j] - mat[i-1][j-1]
			}
		}
	}
	// max := func(a, b int) int {
	// 	if a > b {
	// 		return a
	// 	}
	// 	return b
	// }
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	fmt.Println(mat)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			il := i - K
			jl := j - K
			ir := i + K
			jr := j + K
			a := 0
			b := 0
			c := 0
			d := 0
			if il >= 0 && jl >= 0 {
				a = mat[il][jl]
			}
			if il >= 0 {
				b = mat[il][min(n-1, jr)]
			}
			if jl >= 0 {
				c = mat[min(m-1, ir)][jl]
			}
			d = mat[min(m-1, ir)][min(n-1, jr)]
			ans[i][j] = a + d - b - c
			fmt.Println("i-j", i, "-", j, ans[i][j], il, jl, ir, jr)
		}
	}
	fmt.Println(ans)
	fmt.Println("------------------------")
	return ans
}

// @lc code=end
