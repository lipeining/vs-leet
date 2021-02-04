/*
 * @lc app=leetcode.cn id=747 lang=golang
 *
 * [747] 至少是其他数字两倍的最大数
 */

// @lc code=start
func dominantIndex(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	big := nums[0]
	bigIndex := 0
	for i := 0; i < n; i++ {
		if big < nums[i] {
			big = nums[i]
			bigIndex = i
		}
	}
	for i := 0; i < n; i++ {
		if i != bigIndex {
			if nums[i]*2 > big {
				return -1
			}
		}
	}
	return bigIndex
}

// @lc code=end

