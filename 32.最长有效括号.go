/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */
package main

// func main() {
// 	longestValidParentheses(")()())")
// 	longestValidParentheses(")()()(")
// 	longestValidParentheses(")(()()))")
// 	longestValidParentheses("()(()")
// }

// @lc code=start
func longestValidParentheses(s string) int {
	n := len(s)
	stack := make([]int, 0)
	stack = append(stack, -1)
	ans := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			length := len(stack)
			if length == 0 {
				stack = append(stack, i)
			} else {
				ans = max(ans, i-stack[length-1])
			}
		}
	}
	// fmt.Println(ans)
	return ans
}
func longestValidParenthesesDP(s string) int {
	// n := len(s)
	// dp := make([]bool, n)
	ans := 0
	return ans
}
func match(a, b byte) bool {
	return a == '(' && b == ')'
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
