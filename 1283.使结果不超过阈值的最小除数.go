/*
 * @lc app=leetcode.cn id=1283 lang=golang
 *
 * [1283] 使结果不超过阈值的最小除数
 */

// @lc code=start
func smallestDivisor(nums []int, threshold int) int {
	left := 1
	right := 100000000
	for left < right {
		mid := left + (right-left)>>1
		if helper(nums, threshold, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func helper(nums []int, threshold int, K int) bool {
	sum := 0
	for _, n := range nums {
		sum += div(n, K)
		if sum > threshold {
			return false
		}
	}
	return true
}
func div(a, b int) int {
	n := a / b
	if a%b != 0 {
		return n + 1
	}
	return n
}

// @lc code=end
