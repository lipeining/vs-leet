/*
 * @lc app=leetcode.cn id=456 lang=golang
 *
 * [456] 132模式
 */

// @lc code=start
func find132pattern(nums []int) bool {
	// 3 5 0 3 4 => 3 5 4 true
	ans := false
	stack := make([]int, 0)
	for i:=0;i<len(nums);i++ {
		length := len(stack)
		cur := nums[i]
		if length == 0 {
			stack = append(stack, cur)
		} else if length == 1{
			top := stack[length-1]
			if cur > top {
				stack = append(stack, cur)
			} else {
				stack = stack[0:0]
				stack = append(stack, cur)
			}
		} else if length == 2 {
			top := stack[length-1]
			prev := stack[length-2]
			if prev < cur && cur < top {
				ans = true
				break
			}
			// 保持最多两个节点
			if cur > top {
				stack = append(stack, cur)
				stack = stack[1:]
			} else {
				stack = stack[0:0]
				stack = append(stack, cur)
			}
		}
	}
	return ans
}
// @lc code=end

