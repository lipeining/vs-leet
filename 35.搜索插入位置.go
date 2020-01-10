/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for ; left < right; {
		mid := left + (right-left)/2
		// mid := (left+right)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 此时 left 一定等于 right
	if left < len(nums) && nums[left] < target {
		return left + 1
	}
	return left
}
// @lc code=end

