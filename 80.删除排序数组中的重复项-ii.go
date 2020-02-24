/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	slow, fast := 0, 0
	curLen := 0
	maxLen := 2
	for fast < length {
		for fast < length && nums[slow] == nums[fast] {
			curLen++
			fast++
		}
		if curLen > maxLen {
			curLen = maxLen
		}
		for i := 0; i < curLen; i++ {
			nums[slow+i] = nums[slow]
		}
		slow += curLen
		if fast < length {
			nums[slow] = nums[fast]
		}
		// fmt.Println(nums, slow, fast)
		curLen = 0
	}
	return slow
}

// @lc code=end
