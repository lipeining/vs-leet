/*
 * @lc app=leetcode.cn id=908 lang=golang
 *
 * [908] 最小差值 I
 */

// @lc code=start
func smallestRangeI(nums []int, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	mx, mi := nums[0], nums[0]
	for _, num := range nums {
		mx = max(mx, num)
		mi = min(mi, num)
	}
	return max(0, mx-mi-2*k)
}

// @lc code=end

