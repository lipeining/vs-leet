/*
 * @lc app=leetcode.cn id=921 lang=golang
 *
 * [921] 使括号有效的最少添加
 */

// @lc code=start
func minAddToMakeValid(S string) int {
	stack := make([]string, 0)
	for i:=0;i<len(S);i++ {
		if S[i] == '(' {
			stack = append(stack, "(")
		} else {
			length := len(stack)
			if length == 0 {
				stack = append(stack, ")")
				continue
			}
			top := stack[length-1]
			if top == "(" {
				stack = stack[:length-1]
			} else {
				stack = append(stack, ")")
			}
		}
		fmt.Println(stack)
	}
	fmt.Println(stack)
	return len(stack)
}
// @lc code=end

