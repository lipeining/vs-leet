/*
 * @lc app=leetcode.cn id=801 lang=golang
 *
 * [801] 使序列递增的最小交换次数
 */

// @lc code=start
func minSwap(A []int, B []int) int {
	n := len(A)
	n1, s1 := 0, 1
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// max := func(a, b int) int {
	// 	if a > b {
	// 		return a
	// 	}
	// 	return b
	// }
	// ans := 0
	// n1 表示i-1未被交换的最小交换数
	// s1表示i-1被交换的最小交换数
	for i := 1; i < n; i++ {
		n2 := math.MaxInt32
		s2 := math.MaxInt32
		if A[i-1] < A[i] && B[i-1] < B[i] {
			n2 = min(n2, n1)
			s2 = min(s2, s1+1)
		}
		if A[i-1] < B[i] && B[i-1] < A[i] {
			n2 = min(n2, s1)
			s2 = min(s2, n1+1)
		}
		n1 = n2
		s1 = s2
	}
	ans := min(n1, s1)
	return ans
}

// @lc code=end

