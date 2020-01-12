/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, str string) bool {
	smap := make(map[string]string)
	rmap := make(map[string]string)
	p := strings.Split(pattern, "")
	s := strings.Split(str, " ")
	if len(p) != len(s) {
		return false
	}
	for i:= 0; i < len(p); i++ {
		c := string(p[i])
		m,ok := smap[c]
		if ok {
			if m != s[i] {
				return false
			}
		} else {
			r,rok := rmap[s[i]]
			if rok {
				if r != c {
					return false
				}
			} else {
				rmap[s[i]] = c
			}
			smap[c] = s[i]
		}
	}
	return true
}
// @lc code=end

