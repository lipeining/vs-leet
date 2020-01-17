/*
 * @lc app=leetcode.cn id=709 lang=golang
 *
 * [709] 转换成小写字母
 */

// @lc code=start
func toLowerCase(str string) string {
	ans := ""
	for i:=0;i<len(str);i++ {
		if unicode.IsUpper(rune(str[i])) {
			ans += strings.ToLower(string(str[i]))
		} else {
			ans += string(str[i])
		}
	}
	return ans
}
// @lc code=end

