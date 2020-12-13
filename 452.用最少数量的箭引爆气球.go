/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */
package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}})
// 	findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}})
// 	findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}})
// 	findMinArrowShots([][]int{{1, 2}})
// 	findMinArrowShots([][]int{{1, 2}, {1, 2}})
// 	findMinArrowShots([][]int{{1, 2}, {4, 5}, {1, 5}})
// }

// @lc code=start
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	ans := 1
	pos := points[0][1]
	for _, point := range points {
		if point[0] > pos {
			ans++
			pos = point[1]
		}
	}
	fmt.Println("ans", ans)
	return ans
}

// @lc code=end
