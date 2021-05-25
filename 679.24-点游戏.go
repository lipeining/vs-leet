/*
 * @lc app=leetcode.cn id=679 lang=golang
 *
 * [679] 24 点游戏
 */
package main

import "fmt"

// func main() {
// 	judgePoint24([]int{4, 1, 8, 7})
// 	judgePoint24([]int{1, 2, 1, 2})
// }

// @lc code=start
func judgePoint24(cards []int) bool {
	n := len(cards)
	abs := func(a float64) float64 {
		if a < 0 {
			return -a
		}
		return a
	}
	ADD, MULTIPLY, SUBTRACT, DIVIDE := 0, 1, 2, 3
	var dfs func(val []float64) bool
	dfs = func(val []float64) bool {
		if len(val) == 0 {
			return false
		}
		if len(val) == 1 {
			return abs(val[0]-float64(24)) < 1e-6
		}
		// choose two num
		size := len(val)
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if i == j {
					continue
				}
				left := val[i]
				right := val[j]
				list := make([]float64, 0)
				for k := 0; k < size; k++ {
					if k != i && k != j {
						list = append(list, val[k])
					}
				}
				// +
				for k := 0; k < 4; k++ {
					if k < 2 && i < j {
						continue
					}
					switch k {
					case ADD:
						list = append(list, left+right)
					case MULTIPLY:
						list = append(list, left*right)
					case SUBTRACT:
						list = append(list, left-right)
					case DIVIDE:
						if abs(right) < 1e-6 {
							continue
						} else {
							list = append(list, left/right)
						}
					}
					if dfs(list) {
						return true
					}
					list = list[:len(list)-1]
				}
			}
		}
		return false
	}
	ans := false
	val := make([]float64, n)
	for i := 0; i < n; i++ {
		val[i] = float64(cards[i])
	}
	ans = dfs(val)
	fmt.Println("ansssssssss", ans)
	return ans
}

// @lc code=end
