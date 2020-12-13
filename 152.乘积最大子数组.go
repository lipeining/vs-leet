/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */
package main

// func main() {
// 	maxProduct([]int{2, 3, -2, 4})
// 	maxProduct([]int{-2, 0, -1})
// }

// @lc code=start
func maxProduct(nums []int) int {
	// 保存 min, max 值
	minX, maxX := nums[0], nums[0]
	ans := nums[0]
	for i, num := range nums {
		if i == 0 {
			continue
		}
		a := minX * num
		b := maxX * num
		// fmt.Println(a, b, minX, maxX)
		minX = min(num, min(a, b))
		maxX = max(num, max(a, b))
		ans = max(ans, maxX)
	}
	// fmt.Println("ans, minX, maxX", ans, minX, maxX)
	return ans
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end
