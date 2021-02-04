/*
 * @lc app=leetcode.cn id=1027 lang=golang
 *
 * [1027] 最长等差数列
 */
package main

// func main() {
// 	longestArithSeqLength([]int{3, 6, 9, 12})
// 	longestArithSeqLength([]int{9, 4, 7, 2, 10})
// 	longestArithSeqLength([]int{20, 1, 15, 3, 10, 5, 8})
// }

// @lc code=start
func longestArithSeqLength(A []int) int {
	n := len(A)
	dp := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
	}
	// dp[i] 为，对应下标，各种公差的集合
	// for i := 1; i < n; i++ {
	// 	for j := 0; j < i; j++ {

	// 	}
	// }
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diff := A[j] - A[i]
			if length, ok := dp[i][diff]; ok {
				dp[j][diff] = length + 1
			} else {
				dp[j][diff] = 2
			}
		}
	}
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		for _, v := range dp[i] {
			ans = max(ans, v)
		}
	}
	// fmt.Println("ansssss", ans)
	return ans
}

// @lc code=end
