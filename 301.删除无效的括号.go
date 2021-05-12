/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */

// @lc code=start
func removeInvalidParentheses(s string) []string {
	left, right := 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			right++
		}
	}
}

// @lc code=end

