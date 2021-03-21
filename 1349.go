package main

import "fmt"

// func main() {
// 	// [['#','.','#','#','.','#'],['.','#','#','#','#','.'],['#','.','#','#','.','#']]
// 	// 4
// 	maxStudents([][]byte{
// 		{'#', '.', '#', '#', '.', '#'},
// 		{'.', '#', '#', '#', '#', '.'},
// 		{'#', '.', '#', '#', '.', '#'},
// 	})
// }

// https://leetcode-cn.com/problems/maximum-students-taking-exam/solution/xiang-jie-ya-suo-zhuang-tai-dong-tai-gui-hua-jie-f/
func maxStudents(seats [][]byte) int {
	m := len(seats)
	n := len(seats[0])
	dp := make([][]int, m+1)
	size := 1 << n
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, size)
	}
	hasUser := func(state int, index int) bool {
		return state&(1<<index) == 1
	}
	fmt.Println("hasUser", 0, 1, hasUser(0, 1))
	fmt.Println("hasUser", 1, 1, hasUser(1, 1))
	fmt.Println("hasUser", 2, 1, hasUser(2, 1))
	countOne := func(a int) int {
		c := 0
		for a > 0 {
			a &= a - 1
			c++
		}
		return c
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for row := 1; row <= m; row++ {
		for s := 0; s < size; s++ {
			ok := true
			for k := 0; k < n; k++ {
				if (hasUser(s, k) && seats[row-1][k] == '#') || ((k < n-1) && hasUser(s, k) && hasUser(s, k+1)) {
					ok = false
					break
				}
			}
			if !ok {
				dp[row][s] = -1
				continue
			}
			cs := countOne(s)
			for last := 0; last < size; last++ {
				if dp[row-1][last] == -1 {
					continue
				}
				flag := true
				for k := 0; k < n; k++ {
					if hasUser(last, k) && ((k > 0 && hasUser(s, k-1)) || (k < n-1 && hasUser(s, k+1))) {
						flag = false
						break
					}
				}
				if flag {
					dp[row][s] = max(dp[row][s], dp[row-1][last]+cs)
				}
			}
		}
	}
	fmt.Println(dp)
	ans := 0
	for s := 0; s < size; s++ {
		ans = max(ans, dp[m][s])
	}
	return ans
}
func maxStudentsold(seats [][]byte) int {
	m := len(seats)
	n := len(seats[0])
	dp := make([][]int, m)
	size := 1 << n
	for i := 0; i < m; i++ {
		dp[i] = make([]int, size)
	}
	countOne := func(a int) int {
		c := 0
		for a > 0 {
			a &= a - 1
			c++
		}
		return c
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 以一行作为状态压缩
	for i := 0; i < m; i++ {
		for j := 0; j < size; j++ {
			fix := true
			for k := 0; k < n; k++ {
				if j&(1<<k) == 1 {
					if seats[i][k] == '#' {
						fmt.Println(i, "-", j, "on k", k, "has #")
						fix = false
						break
					} else {
						if (j&(1<<k-1) == 1) || (j&(1<<k+1) == 1) {
							fmt.Println(i, "-", j, "on k", k, "has neibor")
							fix = false
							break
						}
					}
				}
			}
			// 这里需要判断不能相邻坐人
			// 如果本来不能坐人的，那么，直接跳过这个状态，说明这个状态的人数为0
			if !fix {
				continue
			}
			// 这个状态下的人可以坐下，但是需要跟前面一行的状态比对
			cj := countOne(j)
			if i == 0 {
				dp[i][j] = cj
				continue
			}
			prev := dp[i-1]
			for l := 0; l < size; l++ {
				// 如果这个 l  与 j 不冲突，那么，j+=l 的状态数
				pass := true
				for k := 0; k < n; k++ {
					if l&(1<<k) == 1 {
						if (j&(1<<k-1) == 1) || (j&(1<<k+1) == 1) {
							fmt.Println(i, "-", j, "with l", l, "on k", k, "has up-down check")
							pass = false
							break
						}
					}
				}
				if pass {
					dp[i][j] = max(dp[i][j], prev[l]+cj)
				}
			}
		}
		fmt.Println(dp[i])
	}
	fmt.Println(dp)
	ans := 0
	for j := 0; j < size; j++ {
		ans = max(ans, dp[m-1][j])
	}
	return ans
}
