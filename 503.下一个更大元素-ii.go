/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	ans := make([]int, len(nums))
	stack := make([]int, len(nums) * 2)
	for i:=0;i<len(nums);i++ {
		stack[i] = nums[i]
		stack[i+len(nums)] = nums[i]
	}
	for i:=0;i<len(nums);i++ {
		start := i+1
		end := i + len(nums)
		ans[i] = helper(nums[i], stack, start, end)
		// fmt.Print(ans[i])
	}
	// for i:=len(nums)-1;i>=0;i-- {

	// }
	return ans
}
func helper(target int, stack []int, start int, end int) int {
	for i:= start;i<end;i++ {
		if stack[i] > target {
			return stack[i]
		}
	}
	return -1
}
// @lc code=end

