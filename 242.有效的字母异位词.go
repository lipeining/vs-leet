/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	toMap := func(str string) map[byte]int {
		m := make(map[byte]int)
		for i := 0; i < len(str); i++ {
			m[str[i]]++
		}
		return m
	}
	if len(s) != len(t) {
		return false
	}
	sm := toMap(s)
	tm := toMap(t)
	for k, v := range sm {
		if v2, ok := tm[k]; !ok || v2 != v {
			return false
		}
	}
	return true
}

// @lc code=end

