/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	stack := make([]int, 0)
	ans := ""
	n := len(S)
	for i := 0; i < n; i++ {
		if S[i] == '(' {
			stack = append(stack, i)
		} else {
			size := len(stack)
			if size == 1 {
				top := stack[0]
				ans += S[top+1:i]
			}
			stack = stack[:size-1]
		}
	}
	return ans
}
func removeOuterParenthesesOld(S string) string {
	stack := make([]byte, 0)
	ans := ""
	firstFlag := false
	leftCounter := 0
	for i := 0; i < len(S); i++ {
		// length := len(stack)
		if firstFlag {
			if S[i] == ')' && leftCounter == 0 {
				// 这里是右括号
				for j := 0; j < len(stack); j++ {
					ans += string(stack[j])
				}
				stack = stack[0:0]
				firstFlag = false
			} else {
				if S[i] == '(' {
					leftCounter++
				} else {
					leftCounter--
				}
				stack = append(stack, S[i])
			}
		} else {
			// 得到第一个左括号了
			firstFlag = true
		}
	}
	return ans
}

// @lc code=end

