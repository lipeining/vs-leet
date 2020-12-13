/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 */

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	n := len(s1)
	if n == 0 {
		return true
	}
	if s1 == s2 {
		return true
	}
	letters := make([]int, 26)
	for i := 0; i < n; i++ {
		letters[s1[i]-'a']++
		letters[s2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if letters[i] != 0 {
			return false
		}
	}
	dp := make([][][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, n+1)
		}
	}
	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			for j := 0; j <= n-l; j++ {
				if l == 1 {
					dp[i][j][l] = s1[i] == s2[j]
				} else {
					for q := 1; q < l; q++ {
						if dp[i][j][q] && dp[i+q][j+q][l-q] {
							dp[i][j][l] = true
							break
						}
						if dp[i][j+l-q][q] && dp[i+q][j][l-q] {
							dp[i][j][l] = true
							break
						}
					}
				}
			}
		}
	}

	return dp[0][0][n]
}

// @lc code=end

