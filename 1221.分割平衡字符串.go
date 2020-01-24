/*
 * @lc app=leetcode.cn id=1221 lang=golang
 *
 * [1221] 分割平衡字符串
 */

// @lc code=start
func balancedStringSplit(s string) int {
	ans := 0
	l:=0
	for i:=0;i<len(s);i++ {
		if s[i]=='L'{
			l++
		} else {
			l--
		}
		if l == 0 {
			ans++
		}
	}
	return ans
}
// @lc code=end

