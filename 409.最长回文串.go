/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

// @lc code=start
func longestPalindrome(s string) int {
	counter := make(map[rune]int)
	for _,c := range s {
		n,ok := counter[c]
		if ok {
			counter[c] = n + 1
		} else {
			counter[c] = 1
		}
	}
	res := 0
	flag := 0
	for c := range counter {
		if counter[c] == 1 {
			flag = 1
		} else if counter[c] % 2 == 1 {
			res += counter[c] - 1
			flag = 1
		} else {
			res += counter[c]
		}
	}
	return res + flag
}
// @lc code=end

