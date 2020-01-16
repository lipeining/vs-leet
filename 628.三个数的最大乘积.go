/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */
import "sort"
// @lc code=start
func maximumProduct(nums []int) int {
	// 存在负数 需要考虑
	sort.Ints(nums)
	length := len(nums)
	p := nums[length-1] * nums[length-2] * nums[length-3]
	n := nums[length-1] * nums[0] * nums[1]
	if p > n {
		return p
	}
	return n
}
// @lc code=end

