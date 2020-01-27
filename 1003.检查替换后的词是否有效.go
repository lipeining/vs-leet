/*
 * @lc app=leetcode.cn id=1003 lang=golang
 *
 * [1003] 检查替换后的词是否有效
 */

// @lc code=start
func isValid(S string) bool {
	stack := make([]string, 0)
	for i:=0;i<len(S);i++ {
		length := len(stack)
		if length >= 2 && S[i] == 'c' {
			prev := stack[length-2]
			next := stack[length-1]
			if prev == "a" && next == "b" {
				stack = stack[:length-2]
			}
		} else {
			stack = append(stack, string(S[i]))
		}
		fmt.Println(stack)
	}
	return len(stack) == 0
}
// @lc code=end

