/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first,second := math.MaxInt64,math.MaxInt64
	for i:=0;i<len(nums);i++ {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second {
			second = nums[i]
		} else {
			return true
		}
	}
	return false
}
// @lc code=end

