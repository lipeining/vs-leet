/*
 * @lc app=leetcode.cn id=1092 lang=golang
 *
 * [1092] 最短公共超序列
 */

// @lc code=start
func shortestCommonSupersequence(str1 string, str2 string) string {
	// str1 的前半部分和  str2 的后半部分
	// str2 的后半部分和  str2 的前半部分
	// 求其中的最大值，然后就可以得到最短的字符串了

	// 求 LCS ，然后构建答案
	n := len(str1)
	m := len(str2)
	dp := make([][]string, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]string, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + string(str1[i-1])
			} else {
				if len(dp[i-1][j]) > len(dp[i][j-1]) {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	ans := ""
	lcs := dp[n][m]
	i, j := 0, 0
	for _, ch := range lcs {
		for i < n && str1[i] != byte(ch) {
			ans = ans + string(str1[i])
			i++
		}
		for j < m && str2[j] != byte(ch) {
			ans = ans + string(str2[j])
			j++
		}
		ans = ans + string(ch)
		i++
		j++
	}
	return ans + str1[i:] + str2[j:]
}

// @lc code=end

