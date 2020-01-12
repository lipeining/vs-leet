/*
 * @lc app=leetcode.cn id=520 lang=golang
 *
 * [520] 检测大写字母
 */
import "unicode"
// @lc code=start
func detectCapitalUse(word string) bool {
    if word == "" || len(word) == 1 {
		return true
	}
	first := unicode.IsUpper(rune(word[0]))
	second := unicode.IsUpper(rune(word[1]))
	if !first && second {
		return false
	}
	checkUpper := true
	if !second {
		checkUpper = false
	}
	for i:=1;i<len(word);i++ {
		isUpper  := unicode.IsUpper(rune(word[i]))
		if (checkUpper && !isUpper) || (!checkUpper && isUpper) {
			return false
		}
	}
	return true
}
// @lc code=end

