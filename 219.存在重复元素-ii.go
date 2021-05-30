/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	cnt := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		if j, ok := cnt[nums[i]]; ok {
			if i-j <= k {
				return true
			}
		}
		cnt[nums[i]] = i
	}
	return false
}

// @lc code=end

