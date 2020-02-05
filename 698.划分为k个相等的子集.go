import "sort"

/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	// 抄答案
	row := len(nums) - 1
	if nums[row] > target {
		return false
	}
	for row >= 0 && nums[row] == target {
		row--
		k--
	}
	groups := make([]int, k)
	return search(groups, row, nums, target)
}
func search(groups []int, row int, nums []int, target int) bool {
	if row < 0 {
		return true
	}
	v := nums[row]
	row--
	for i := 0; i < len(groups); i++ {
		if groups[i]+v <= target {
			groups[i] += v
			if search(groups, row, nums, target) {
				return true
			}
			groups[i] -= v
		}
		if groups[i] == 0 {
			break
		}
	}
	return false
}

// @lc code=end
