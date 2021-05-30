/*
 * @lc app=leetcode.cn id=1023 lang=golang
 *
 * [1023] 驼峰式匹配
 */
package main

import "fmt"

func main() {
	// camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB")
	// camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBa")
	camelMatch([]string{"IXfGawluvnCa", "IsXfGaxwulCa", "IXfGawlqtCva", "IXjfGawlmeCa", "IXfGnaynwlCa", "IXfGawlCa"}, "IXfGawlCa")
}

// @lc code=start
func camelMatch(queries []string, pattern string) []bool {
	length := len(queries)
	ans := make([]bool, length)
	check := func(left, right string) bool {
		n := len(left)
		m := len(right)
		i := 0
		j := 0
		for i < n && j < m {
			if left[i] == right[j] {
				j++
				i++
			} else {
				if left[i] >= 'A' && left[i] <= 'Z' {
					// fmt.Println("on i", left, i, right, j)
					return false
				}
				i++
			}
		}
		if i == n && j == m {
			return true
		}
		if i == n && j != m {
			// fmt.Println("on n", left, i)
			return false
		}
		if j == m {
			// 保证之后都是小写
			for i < n {
				if left[i] >= 'A' && left[i] <= 'Z' {
					// fmt.Println("on i", left, i)
					return false
				}
				i++
			}
		}
		return true
	}
	for i := 0; i < length; i++ {
		ans[i] = check(queries[i], pattern)
	}
	fmt.Println("anssssss", ans)
	return ans
}

// @lc code=end
