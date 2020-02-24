/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] 最大连续1的个数 III
 */

// @lc code=start
func longestOnes(A []int, K int) int {
	left, right, result, count := 0, 0, 0, 0
	for right < len(A) {
		if A[right] == 0 {
			count++
		}
		for count > K {
			if A[left] == 0 {
				count--
			}
			left++
		}
		result = max(result, right-left+1)
		right++
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
