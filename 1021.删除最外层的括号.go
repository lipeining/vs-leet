/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	stack := make([]byte, 0)
	ans:=""
	firstFlag := false
	leftCounter := 0
	for i:=0; i < len(S); i++ {
		// length := len(stack)
		if firstFlag {
			if S[i] == ')'  && leftCounter == 0{
				// 这里是右括号
				for j:=0;j<len(stack);j++ {
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

