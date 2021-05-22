/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 */

// @lc code=start
func buddyStrings(a string, b string) bool {
	al := len(a)
	bl := len(b)
	if al != bl {
		return false
	}
	a1, a2 := -1, -1
	m := make(map[byte]int)
	for i := 0; i < al; i++ {
		if a[i] != b[i] {
			if a1 == -1 {
				a1 = i
			} else {
				if a2 == -1 {
					a2 = i
				} else {
					return false
				}
			}
		}
		m[a[i]]++
	}
	if a1 == -1 && a2 == -1 {
		for _, v := range m {
			if v >= 2 {
				return true
			}
		}
		return false
	}
	if a1 != -1 && a2 == -1 {
		return false
	}
	if a[a2] == b[a1] && a[a1] == b[a2] {
		return true
	}
	return false
}

// @lc code=end

