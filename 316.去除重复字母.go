/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 */
package main

import "fmt"

// func main() {
// 	removeDuplicateLetters("cbacdcbc")
// }

// @lc code=start
func removeDuplicateLetters(s string) string {
	n := len(s)
	cnt := make([]int, 26)
	for i := 0; i < n; i++ {
		cnt[s[i]-'a']++
	}
	ans := ""
	stack := make([]byte, 0)
	instack := func(c byte) bool {
		for i := 0; i < len(stack); i++ {
			if stack[i] == c {
				return true
			}
		}
		return false
	}
	for i := 0; i < n; i++ {
		c := s[i]
		if !instack(c) {
			for len(stack) != 0 && c < stack[len(stack)-1] && cnt[stack[len(stack)-1]-'a'] > 0 {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, c)
		}
		cnt[c-'a']--
	}
	for i := 0; i < len(stack); i++ {
		ans += string(stack[i])
	}
	fmt.Println("anssssssssss", ans)
	return ans
}

// @lc code=end
