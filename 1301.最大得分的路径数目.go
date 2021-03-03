/*
 * @lc app=leetcode.cn id=1301 lang=golang
 *
 * [1301] 最大得分的路径数目
 */
package main

import "fmt"

// func main() {
// 	pathsWithMaxScore([]string{"E23", "2X2", "12S"})
// 	pathsWithMaxScore([]string{"E12", "1X1", "21S"})
// 	pathsWithMaxScore([]string{"E11", "XXX", "11S"})
// 	pathsWithMaxScore([]string{"E11", "X1X", "11S"})
// 	pathsWithMaxScore([]string{"E11345","X452XX","3X43X4","44X312","23452X","1342XS"})
// }

// @lc code=start
func pathsWithMaxScore(board []string) []int {
	n := len(board)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	dp[n-1][n-1][0] = 0
	dp[n-1][n-1][1] = 1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 0 最大值 1 方案数
	mod := int(1e9 + 7)
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == n-1 && j == n-1 {
				dp[i][j][0] = 0
				dp[i][j][1] = 1
			} else if i == n-1 {
				if board[i][j] == 'X' {
					dp[i][j][0] = 0
					dp[i][j][1] = 0
				} else {
					num := int(board[i][j] - '0')
					dp[i][j][0] = dp[i][j+1][0] + num
					if dp[i][j+1][1] == 0 {
						dp[i][j][0] = 0
					}
					dp[i][j][1] = dp[i][j+1][1]
				}
			} else if j == n-1 {
				if board[i][j] == 'X' {
					dp[i][j][0] = 0
					dp[i][j][1] = 0
				} else {
					num := int(board[i][j] - '0')
					dp[i][j][0] = dp[i+1][j][0] + num
					if dp[i+1][j][1] == 0 {
						dp[i][j][0] = 0
					}
					dp[i][j][1] = dp[i+1][j][1]
				}
			} else {
				if board[i][j] == 'X' {
					dp[i][j][0] = 0
					dp[i][j][1] = 0
				} else {
					num := int(board[i][j] - '0')
					if i == 0 && j == 0 {
						// continue
						num = 0
					}
					// 计算 左，上，左上方 三条路线的值
					left := dp[i][j+1][0]
					up := dp[i+1][j][0]
					tri := dp[i+1][j+1][0]
					maybe := max(left, max(up, tri))
					dp[i][j][0] = maybe + num
					// noways
					noways := dp[i][j+1][1] + dp[i+1][j][1] + dp[i+1][j+1][1]
					if noways == 0 {
						dp[i][j][0] = 0
					}
					if left == maybe {
						dp[i][j][1] = (dp[i][j][1] + dp[i][j+1][1]) % mod
					}
					if up == maybe {
						dp[i][j][1] = (dp[i][j][1] + dp[i+1][j][1]) % mod
					}
					if tri == maybe {
						dp[i][j][1] = (dp[i][j][1] + dp[i+1][j+1][1]) % mod
					}
				}
			}
		}
	}
	ans := make([]int, 2)
	ans[0] = dp[0][0][0]
	ans[1] = dp[0][0][1]
	fmt.Println("ansssss", ans)
	return ans
}

// @lc code=end
