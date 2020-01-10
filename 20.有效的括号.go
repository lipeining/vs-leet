/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]string, 0)
	smap := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
    for i:=0; i < len(s); i++ {
		tmp, ok := smap[string(s[i])]
		if !ok {
			stack = append(stack, string(s[i]))
		} else {
			size := len(stack)
			if size == 0 {
				return false
			} else {
				top := stack[size-1]
				if top != tmp {
					return false
				} else {
					stack = stack[:size-1]
				}
			}
		}
	}
	return len(stack) == 0
}
// @lc code=end

