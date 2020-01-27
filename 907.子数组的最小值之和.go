/*
 * @lc app=leetcode.cn id=907 lang=golang
 *
 * [907] 子数组的最小值之和
 */

// @lc code=start
func sumSubarrayMins(A []int) int {
	toMod := 1000000007
	res := 0

	nums := A
	if len(nums) == 0 {
		return 0
	}
	for i:=0;i<len(nums);i++ {
		for j:=i+1;j<=len(nums);j++ {
			sub := nums[i:j]
			res += helper(sub)
		}
	}
	return res % toMod
}
func helper(nums []int) int {
	min := nums[0]
	for i:=1; i<len(nums);i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
// @lc code=end

