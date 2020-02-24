/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	origin := getMap(s1)
	len1 := len(s1)
	window := getMap(s2[:len1])
	right := 0
	length := len(s2)

	for right < length-len1 {
		if check(origin, window) {
			return true
		}
		char := string(s2[right+len1])
		window[char]++
		window[string(s2[right])]--
		right++
	}
	return check(origin, window)
}
func check(origin, window map[string]int) bool {
	for k1, v1 := range origin {
		v2, ok := window[k1]
		if !ok {
			return false
		} else if v1 != v2 {
			return false
		}
	}
	return true
}
func getMap(str string) map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(str); i++ {
		m[string(str[i])]++
	}
	return m
}

// @lc code=end
