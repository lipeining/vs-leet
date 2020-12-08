/*
 * @lc app=leetcode.cn id=467 lang=golang
 *
 * [467] 环绕字符串中唯一的子字符串
 */
// package main

// import "fmt"

// func main() {
// 	findSubstringInWraproundString("a")
// 	findSubstringInWraproundString("cac")
// 	findSubstringInWraproundString("zab")
// }

// @lc code=start
func findSubstringInWraproundString(p string) int {
	n := len(p)
	if n <= 1 {
		return n
	}
	p = "^" + p
	ans := 0
	cnt := 1
	dp := make(map[byte]int)
	for j := 1; j <= n; j++ {
		if p[j]-p[j-1] == 1 || p[j-1] == 'z' && p[j] == 'a' {
			cnt += 1
		} else {
			cnt = 1
		}
		dp[p[j]] = max(dp[p[j]], cnt)
	}
	for _, num := range dp {
		ans += num
	}
	fmt.Println("ans", ans)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
