package main

import "fmt"

func main() {
	getMaximumXor([]int{0, 1, 1, 3}, 2)
	getMaximumXor([]int{2, 3, 4, 7}, 3)
	getMaximumXor([]int{0, 1, 2, 2, 5, 7}, 3)
}
func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	ans := make([]int, n)
	total := nums[0]
	for i := 1; i < n; i++ {
		total ^= nums[i]
	}
	// fmt.Println("init total", total)
	big := 1<<maximumBit - 1
	for i := 0; i < n; i++ {
		last := nums[n-i-1]
		// fmt.Println("notLast", notLastTotal, "total", total, "last", last)
		ans[i] = big ^ total
		total = total ^ last
	}
	fmt.Println("anssss", ans)
	return ans
}
func countPoints(points [][]int, queries [][]int) []int {
	n := len(queries)
	in := func(index int) int {
		cnt := 0
		for j := 0; j < len(points); j++ {
			// 直径
			x, y, r := queries[index][0], queries[index][1], queries[index][2]
			x1, y1 := points[j][0], points[j][1]
			dx, dy := x1-x, y1-y
			if dx*dx+dy*dy <= r*r {
				cnt++
			}
		}
		return cnt
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = in(i)
	}
	return ans
}
