/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */
package main

// func main() {
// 	nextGreaterElements([]int{1, 2, 1}) // 2 -1 2
// }

// @lc code=start
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	stack := make([]int, 0)
	for i := 0; i <= 2*n-2; i++ {
		index := i % n
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			topIndex := top % n
			if nums[topIndex] >= nums[index] {
				break
			}
			stack = stack[:len(stack)-1]
			ans[topIndex] = nums[index]
		}
		stack = append(stack, i)
		// fmt.Println("i", i, "index", index, "ans, stack", ans[index], stack)
	}
	// fmt.Println("anssssss", ans)
	return ans
}
func nextGreaterElements2(nums []int) []int {
	ans := make([]int, len(nums))
	stack := make([]int, len(nums)*2)
	for i := 0; i < len(nums); i++ {
		stack[i] = nums[i]
		stack[i+len(nums)] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		start := i + 1
		end := i + len(nums)
		ans[i] = helper(nums[i], stack, start, end)
		// fmt.Print(ans[i])
	}
	// for i:=len(nums)-1;i>=0;i-- {

	// }
	return ans
}
func helper(target int, stack []int, start int, end int) int {
	for i := start; i < end; i++ {
		if stack[i] > target {
			return stack[i]
		}
	}
	return -1
}

// @lc code=end
