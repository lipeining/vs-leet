import "fmt"

/*
 * @lc app=leetcode.cn id=845 lang=golang
 *
 * [845] 数组中的最长山脉
 */

// @lc code=start
func longestMountain(A []int) int {
	length := len(A)
	ans := 0
	for i := 1; i < length-1; i++ {
		tmp := expand(A, i)
		ans = max(ans, tmp)
	}
	return ans
}
func expand(A []int, index int) int {
	left := index
	right := index
	length := len(A)

	for left >= 1 {
		if A[left] <= A[left-1] {
			break
		}
		left--
	}
	for right < length-1 {
		if A[right] <= A[right+1] {
			break
		}
		right++
	}
	if right == index || left == index {
		return 0
	}
	if right-left < 2 {
		return 0
	}
	return right - left + 1
}

// func longestMountainDP(A []int) int {
// 	// dp[i][j] true false  ans = max(ans, j-i+1)
// 	length := len(A)
// 	dp := make([][]bool, length)
// 	for i := 0; i < length; i++ {
// 		dp[i] = make([]bool, length)
// 	}
// 	ans := 0
// 	for i := 0; i < length-2; i++ {
// 		for j := i + 2; j < length; j++ {
// 			if j-i == 2 {
// 				// check true
// 				if A[i] < A[i+1] && A[i+1] > A[j] {
// 					dp[i][j] = true
// 				}
// 			} else {
// 				// check i down check j down
// 				if A[i] < A[i+1] && A[j] < A[j-1] {
// 					dp[i][j] = dp[i+1][j-1]
// 				}
// 			}
// 			fmt.Println(i, j, dp[i][j])
// 			if dp[i][j] {
// 				ans = max(ans, j-i+1)
// 			}
// 		}
// 	}
// 	return ans
// }
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
