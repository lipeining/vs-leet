/*
 * @lc app=leetcode.cn id=1190 lang=golang
 *
 * [1190] 反转每对括号间的子串
 */

// @lc code=start
func reverseParentheses(s string) string {
	stack := make([]string, 0)
	for i:=0;i<len(s);i++ {
		if s[i] == '(' {
			stack = append(stack, string(s[i]))
		} else if s[i] == ')' {
			length := len(stack)
			j:=length-1
			str := ""
			for j>=0 {
				if stack[j] != "(" {
					str += stack[j]
				} else {
					break
				}
				j--
			}
			stack = stack[:j]
			for k:=0;k<len(str);k++ {
				stack = append(stack, string(str[k]))
			}
		} else {
			stack = append(stack, string(s[i]))
		}
		// fmt.Println(stack)
	}
	return strings.Join(stack, "")
}
// @lc code=end

