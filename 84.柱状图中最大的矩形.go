/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */
package main

import "fmt"

// func main() {
// 	largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
// 	largestRectangleArea([]int{1, 1})
// 	largestRectangleArea([]int{2, 1, 2})
// }

// @lc code=start
func largestRectangleArea(heights []int) int {
	ans := 0
	stack := make([]int, 0)
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	n := len(heights)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		for len(stack) != 0 && heights[stack[len(stack)-1]] > heights[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := 1
			if len(stack) != 0 {
				width = i - stack[len(stack)-1] - 1
			} else {
				width = i - top
			}
			// fmt.Println(ans, "top", top, "width", width, "i", i)
			ans = max(ans, width*heights[top])
		}
		stack = append(stack, i)
	}
	fmt.Println("ans", ans)
	return ans
}

// @lc code=end
