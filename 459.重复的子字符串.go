/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	length := len(s)
    for i:= 1; i<=length/2;i++ {
		if length % i == 0 {
			if helper(s, i) {
				return true
			}
		}
	}
	return false
}
func helper(s string, n int) bool {
	for i:=0; i < n; i++ {
		for j:= 1; j < len(s)/n; j++ {
			if s[i] != s[i+n*j] {
				return false
			}
		}
	}
	return true
}
// @lc code=end

