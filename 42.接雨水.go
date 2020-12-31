/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */
package main

import "fmt"

// func main() {
// 	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
// 	trap([]int{4, 2, 3})
// }

// @lc code=start
func trap(height []int) int {
	// 递增栈
	n := len(height)
	i := 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i < n {
		if height[i] == 0 {
			i++
		} else {
			break
		}
	}
	ans := 0
	// 左边到右边一次
	// 右边到左边一次
	// 都是单调递增栈
	last := i
	for i < n {
		if last != i && height[last] <= height[i] {
			// 计算 last 到 i 之间的雨水
			for j := last + 1; j != i; j++ {
				jValue := min(height[last], height[i]) - height[j]
				// fmt.Println("j", j, "jValue", jValue)
				ans += jValue
			}
			last = i
		}
		i++
	}
	i = n - 1
	for i > last {
		if height[i] == 0 {
			i--
		} else {
			break
		}
	}
	// fmt.Println("last", last, "i", i)
	prev := i
	for i >= last {
		if prev != i && height[prev] <= height[i] {
			for j := i + 1; j != prev; j++ {
				jValue := min(height[prev], height[i]) - height[j]
				// fmt.Println("right j", j, "jValue", jValue)
				ans += jValue
			}
			prev = i
		}
		i--
	}
	fmt.Println("ans----------------", ans)
	return ans
}

// @lc code=end
