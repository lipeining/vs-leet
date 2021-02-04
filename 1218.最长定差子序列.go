/*
 * @lc app=leetcode.cn id=1218 lang=golang
 *
 * [1218] 最长定差子序列
 */
package main

// import "fmt"

// func main() {
// 	longestSubsequence([]int{1, 2, 3, 4}, 1)
// 	longestSubsequence([]int{1, 3, 5, 7}, 1)
// 	longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2)
// }

// @lc code=start
func longestSubsequence(arr []int, difference int) int {
	n := len(arr)
	ans := 1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		t := m[arr[i]-difference] + 1
		m[arr[i]] = t
		ans = max(ans, t)
	}
	return ans
	// n := len(arr)
	// dp := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = 1
	// }
	// for i := 0; i < n; i++ {
	// 	for j := i + 1; j < n; j++ {
	// 		diff := arr[j] - arr[i]
	// 		if diff == difference {
	// 			dp[j] = dp[i] + 1
	// 		}
	// 	}
	// }
	// ans := 0
	// max := func(a, b int) int {
	// 	if a > b {
	// 		return a
	// 	}
	// 	return b
	// }
	// for i := 0; i < n; i++ {
	// 	ans = max(ans, dp[i])
	// }
	// // fmt.Println("ansssss", ans)
	// return ans
}

// @lc code=end
