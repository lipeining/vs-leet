/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	ans := make([]int, 0)
	for _, num := range nums {
		abs := num
		if abs < 0 {
			abs = -abs
		}
		if nums[abs-1] < 0 {
			ans = append(ans, abs)
		} else {
			nums[abs-1] = -nums[abs-1]
		}
	}
	return ans
}
// @lc code=end

