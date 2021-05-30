/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	n := len(s)
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := 0; i < n; i++ {
		l := s[i]
		r := t[i]
		if to, ok := s2t[l]; !ok {
			// 不同字符不能映射到同一个字符上
			if _, ok2 := t2s[r]; ok2 {
				return false
			}
			// 每个出现的字符都应当映射到另一个字符
			s2t[l] = r
			t2s[r] = l
		} else {
			if to != r {
				return false
			}
		}
	}
	return true
}

// @lc code=end

