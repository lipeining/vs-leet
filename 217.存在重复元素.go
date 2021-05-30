/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	cnt := make(map[int]bool)
	n := len(nums)
	for i := 0; i < n; i++ {
		if _, ok := cnt[nums[i]]; ok {
			return true
		}
		cnt[nums[i]] = true
	}
	return false
}

// @lc code=end

