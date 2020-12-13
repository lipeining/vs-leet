/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */
package main

import "fmt"

// func main() {
// 	// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
// 	// 输出：true
// 	// 示例 2：

// 	// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
// 	// 输出：false
// 	// 示例 3：

// 	// 输入：s1 = "", s2 = "", s3 = ""
// 	// 输出：true
// 	isInterleave("aabcc", "dbbca", "aadbbcbcac")
// 	isInterleave("aabcc", "dbbca", "aadbbbaccc")
// 	// isInterleave("", "", "")
// }

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)
	z := len(s3)
	if n+m != z {
		return false
	}
	if z == 0 {
		return true
	}
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j
			if i > 0 {
				if s1[i-1] == s3[p-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
			if j > 0 {
				if s2[j-1] == s3[p-1] {
					dp[i][j] = dp[i][j] || dp[i][j-1]
				}
			}
		}
	}
	// fmt.Println(dp[n][m])
	// fmt.Println("--------------------------")
	return dp[n][m]
}
func isInterleaveZ(s1 string, s2 string, s3 string) bool {
	// dp[i][j][k] 表示 s1[i] s2[j] s3[k] 是否合法
	// 分情况讨论
	// 1. s1[i-1] == s3[k-1] && dp[i-1][j][k-1]
	// 2. s2[j-1] == s3[k-1] && dp[i][j-1][k-1]

	// 只要其中一种成立即可

	n := len(s1)
	m := len(s2)
	z := len(s3)
	if n+m != z {
		return false
	}
	if z == 0 {
		return true
	}
	dp := make([][][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]bool, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]bool, z+1)
		}
	}
	for i := 1; i <= n; i++ {
		if s1[i-1] == s3[i-1] {
			dp[i][0][i] = true
		} else {
			break
		}
	}
	for j := 1; j <= m; j++ {
		if s2[j-1] == s3[j-1] {
			dp[0][j][j] = true
		} else {
			break
		}
	}
	dp[0][0][0] = true
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := 1; k <= i+j; k++ {
				if s1[i-1] == s3[k-1] {
					dp[i][j][k] = dp[i][j][k] || dp[i-1][j][k-1]
				}
				if s2[j-1] == s3[k-1] {
					dp[i][j][k] = dp[i][j][k] || dp[i][j-1][k-1]
				}
				// fmt.Println("i,j,k,ans", i, j, k, dp[i][j][k])
			}
		}
	}
	fmt.Println(dp[n][m][z])
	fmt.Println("--------------------------")
	return dp[n][m][z]
}

// @lc code=end
