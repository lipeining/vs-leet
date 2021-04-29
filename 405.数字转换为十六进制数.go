/*
 * @lc app=leetcode.cn id=405 lang=golang
 *
 * [405] 数字转换为十六进制数
 */

// @lc code=start
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	s := "0123456789abcdef"

	num &= 0xFFFFFFFF
	ans := ""
	for num > 0 {
		ans = string(s[num&0xF]) + ans
		num >>= 4
	}
	return ans
	// ans := make([]string, 0)
	// for num > 0 {
	// 	ans = append(ans, string(s[num&0xF]))
	// 	num >>= 4
	// }
	// ret := ""
	// for i := len(ans) - 1; i >= 0; i-- {
	// 	ret += ans[i]
	// }
	// return ret
}

// @lc code=end

