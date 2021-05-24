/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	n := len(s)
	l := 0
	r := 0
	tb := make(map[byte]int)
	for _, char := range t {
		tb[byte(char)]++
	}
	sb := make(map[byte]int)
	ans := ""
	check := func() bool {
		if len(sb) < len(tb) {
			return false
		}
		for k, v := range tb {
			if v2, ok := sb[k]; !ok || v2 < v {
				return false
			}
		}
		return true
	}
	for r < n {
		sb[s[r]]++
		for check() {
			now := s[l : r+1]
			if ans == "" || len(now) < len(ans) {
				ans = now
			}
			sb[s[l]]--
			l++
		}
		r++
	}
	return ans
}

// @lc code=end

