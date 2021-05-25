/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	slow, fast := 0, 0
	m := make(map[string]int)
	ans := 0
	length := len(s)
	for fast != length {
		char := string(s[fast])
		if _, ok := m[char]; ok {
			curLen := len(m)
			if ans < curLen {
				ans = curLen
			}
			delete(m, string(s[slow]))
			slow++
		} else {
			m[char] = fast
			fast++
		}
	}
	if len(m) > ans {
		ans = len(m)
	}
	return ans
}

// @lc code=end
