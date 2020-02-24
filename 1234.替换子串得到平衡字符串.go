/*
 * @lc app=leetcode.cn id=1234 lang=golang
 *
 * [1234] 替换子串得到平衡字符串
 */

// @lc code=start
func balancedString(s string) int {
	n := len(s)
	freq := make(map[string]int)
	for _, c := range s {
		freq[string(c)]++
	}
	limit := n / 4
	if check(freq, limit) {
		return 0
	}
	j := 0
	ans := n
	for i := 0; i < n; i++ {
		freq[string(s[i])]--
		for check(freq, limit) {
			ans = min(ans, i-j+1)
			freq[string(s[j])]++
			j++
		}
	}
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func check(freq map[string]int, limit int) bool {
	for _, v := range freq {
		if v > limit {
			return false
		}
	}
	return true
}

// @lc code=end
