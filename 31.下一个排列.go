/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	n := len(nums)
	for r := n - 1; r > 0; r-- {
		l := r - 1
		for ; l >= 0; l-- {
			if nums[r] > nums[l] {
				break
			}
		}
		if l >= 0 {
			nums[l], nums[r] = nums[r], nums[l]
			return
		}
	}
	sort.Ints(nums)
}

// @lc code=end

