/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */
package main

import "fmt"

// func main() {
// 	decodeString("3[a]2[bc]")
// 	decodeString("3[a2[c]]")
// 	decodeString("2[abc]3[cd]ef")
// 	decodeString("abc3[cd]xyz")
// }

// @lc code=start
func decodeString(s string) string {
	n := len(s)
	i := 0
	str := make([]string, 0)
	mul := make([]int, 0)
	multi := 0
	res := ""
	for i < n {
		if s[i] == '[' {
			str = append(str, res)
			mul = append(mul, multi)
			multi = 0
			res = ""
		} else if s[i] == ']' {
			times := mul[len(mul)-1]
			mul = mul[:len(mul)-1]
			curStr := ""
			for i := 0; i < times; i++ {
				curStr += res
			}
			res = str[len(str)-1] + curStr
			str = str[:len(str)-1]
		} else if s[i] >= '0' && s[i] <= '9' {
			multi = multi*10 + int(s[i]-'0')
		} else {
			res += string(s[i])
		}
		i++
	}
	fmt.Println("anssss", res)
	return res
}

// @lc code=end
