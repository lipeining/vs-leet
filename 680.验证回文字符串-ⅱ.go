/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 */

// @lc code=start
func validPalindrome(s string) bool {
	// 双指针是应该的吧
	l:=0
	length:=len(s)
	r:=length-1
	for l < r {
		if s[l] != s[r] {
			// 左右挑选一个
			return helper(s[l+1:r+1]) || helper(s[l:r])
		} else {
			l++
			r--
		}
	}
	return true
}
func helper(s string) bool {
	fmt.Println(s)
	l:=0
	length:=len(s)
	r:=length-1
	for l < r {
		if s[l] != s[r] {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}
// @lc code=end

