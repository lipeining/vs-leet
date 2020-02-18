/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	if length == 1 {
		if nums[0] == target {
			return true
		}
		return false
	}
	index := findIndex(nums)
	if nums[index] == target {
		return true
	}
	if index == 0 {
		return bins(nums, 0, length-1, target)
	}
	if nums[0] > target {
		return bins(nums, index, length-1, target)
	}
	return bins(nums, 0, index-1, target)
}
func bins(nums []int, start, end, target int) bool {
	for start <= end {
		mid := start + (end-start)>>1
		if nums[mid] == target {
			return true
		} else {
			if nums[mid] > target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return false
}
func findIndex(nums []int) int {
	left := 0
	right := len(nums) - 1
	if nums[left] < nums[right] {
		return 0
	}
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[mid+1] {
			return mid + 1
		} else {
			if nums[mid] < nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return 0
}
// @lc code=end

