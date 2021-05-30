/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	cnt := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++ {
		cnt[s[i]]++
	}
	for i := 0; i < n; i++ {
		if cnt[s[i]] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end

