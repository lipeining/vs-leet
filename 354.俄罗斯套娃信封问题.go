/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(a, b int) bool {
		return envelopes[a][0] < envelopes[b][0]
	})
	n := len(envelopes)
	dp := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i]+1)
	}
	// fmt.Println(envelopes, dp, ans)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

