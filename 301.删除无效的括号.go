/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */
package main

import "fmt"

// func main() {
// 	removeInvalidParentheses("()())()")
// 	removeInvalidParentheses("(a)())()")
// 	removeInvalidParentheses(")(")
// }

// @lc code=start
func removeInvalidParentheses(s string) []string {
	left, right := 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	// 最终的到需要删除的 left right

	ans := make([]string, 0)
	ansMap := make(map[string]bool)
	var dfs func(lc, rc, le, re int, index int, path string)
	dfs = func(lc, rc, le, re int, index int, path string) {
		if index == n {
			if le == left && re == right {
				// ans = append(ans, path)
				ansMap[path] = true
			}
			return
		}
		// 只有左括号严格大于右括号的数量时
		// lc>rc 才可以添加 右括号
		if s[index] == '(' {
			if le < left {
				dfs(lc, rc, le+1, re, index+1, path)
			}
			dfs(lc+1, rc, le, re, index+1, path+"(")
		} else if s[index] == ')' {
			if re < right {
				dfs(lc, rc, le, re+1, index+1, path)
			}
			if lc > rc {
				dfs(lc, rc+1, le, re, index+1, path+")")
			}
		} else {
			dfs(lc, rc, le, re, index+1, path+string(s[index]))
		}
	}
	dfs(0, 0, 0, 0, 0, "")
	for k := range ansMap {
		ans = append(ans, k)
	}
	fmt.Println("l-r", left, "-", right)
	// fmt.Println("anssssssssss", ans)
	return ans
}

// @lc code=end
