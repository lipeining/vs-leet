/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	if nums[left] < nums[right] {
		return nums[left]
	}
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		} else {
			if nums[mid] < nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return nums[left]
}

// @lc code=end
