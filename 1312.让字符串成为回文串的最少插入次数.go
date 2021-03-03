/*
 * @lc app=leetcode.cn id=1312 lang=golang
 *
 * [1312] 让字符串成为回文串的最少插入次数
 */
package main

// import "fmt"

// func main() {
// 	minInsertions("zzazz")
// 	minInsertions("mbadm")
// 	minInsertions("leetcode")
// 	minInsertions("g")
// 	minInsertions("no")
// }

// @lc code=start
func minInsertions(s string) int {
	n := len(s)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := i; j < n; j++ {
			memo[i][j] = -1
		}
		memo[i][i] = 0
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		need := memo[i][j]
		if need != -1 {
			// fmt.Println("i-j ", i, "-", j, need)
			return need
		}
		if s[i] == s[j] {
			next := dfs(i+1, j-1)
			memo[i][j] = next
			// fmt.Println("i-j ", i, "-", j, next)
			return next
		}
		// 求最小的方案 + 1
		// 左边添加或者右边添加
		left := dfs(i+1, j)
		right := dfs(i, j-1)
		res := left
		if left > right {
			res = right
		}
		res++
		memo[i][j] = res
		// fmt.Println("i-j ", i, "-", j, left, right, res)
		return res
	}
	ans := dfs(0, n-1)
	// fmt.Println("ansssss", ans)
	return ans
}

// @lc code=end
