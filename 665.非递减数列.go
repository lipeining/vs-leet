/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */

// @lc code=start
func checkPossibility(nums []int) bool {
	length := len(nums)
	if length <= 1 {
		return true
	}
	l:=0
	r:=length-1
	for l < length - 1 && r > 0 {
		if nums[l] <= nums[l+1] || nums[r-1] <= nums[r] {
			if nums[l] <= nums[l+1] {
				l++
			}
			if nums[r-1] <= nums[r] {
				r--
			}
		} else {
			break
		}
	}
	if l < r - 1 {
		return false
	} else if l >=r {
		return true
	} else if l == 0 || r == length-1 {
		return true
	}
	start := l-1
	end := r+1
	// 换 l 
	// 换 r
	return nums[start] <= nums[r] || nums[l] <= nums[end]
}
// @lc code=end

