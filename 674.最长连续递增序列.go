/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	// 快慢指针加 max 定值
	max := 0
	length := len(nums)
	i:=0;
	for i<length {
		j:=i;
		for j<length-1 {
			if nums[j]>=nums[j+1] {
				break
			}
			j++
		}
		len := j-i+1
		if len > max {
			max = len
		}
		i = j+1
	}
	return max
}
// @lc code=end

