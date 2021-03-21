package main

import (
	"fmt"
	"sort"
)

func main() {
	func1()
	func2()
	func3()
	func4()
	fmt.Println("anssssssssssssssss end")
}
func func1() {}
func countGoodRectangles(rectangles [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n := len(rectangles)
	counter := make(map[int]int)
	maxLen := 0
	for i := 0; i < n; i++ {
		li := rectangles[i][0]
		wi := rectangles[i][1]
		length := min(li, wi)
		maxLen = max(maxLen, length)
		counter[length]++
	}
	ans := counter[maxLen]
	return ans
}
func func2() {
	// 	输入：nums = [2,3,4,6]
	// 输出：8
	// 解释：存在 8 个满足题意的元组：
	// (2,6,3,4) , (2,6,4,3) , (6,2,3,4) , (6,2,4,3)
	// (3,4,2,6) , (3,4,2,6) , (3,4,6,2) , (4,3,6,2)
	// 示例 2：

	// 输入：nums = [1,2,4,5,10]
	// 输出：16
	// 解释：存在 16 个满足题意的元组：
	// (1,10,2,5) , (1,10,5,2) , (10,1,2,5) , (10,1,5,2)
	// (2,5,1,10) , (2,5,10,1) , (5,2,1,10) , (5,2,10,1)
	// (2,10,4,5) , (2,10,5,4) , (10,2,4,5) , (10,2,4,5)
	// (4,5,2,10) , (4,5,10,2) , (5,4,2,10) , (5,4,10,2)
	// 示例 3：

	// 输入：nums = [2,3,4,6,8,12]
	// 输出：40
	// 示例 4：

	// 输入：nums = [2,3,5,7]
	// 输出：0
	tupleSameProduct([]int{2, 3, 4, 6})
	tupleSameProduct([]int{1, 2, 4, 5, 10})
	tupleSameProduct([]int{2, 3, 4, 6, 8, 12})
	tupleSameProduct([]int{2, 3, 5, 7})
}
func tupleSameProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	// a, b , c, d
	ans := 0
	counter := make(map[int][][]int)
	for i := 0; i < n; i++ {
		a := nums[i]
		for j := i + 1; j < n; j++ {
			b := nums[j]
			m := a * b
			counter[m] = append(counter[m], []int{a, b})
		}
	}
	// fmt.Println(counter)
	for _, v := range counter {
		size := len(v)
		if size >= 2 {
			ans += 8 * (size * (size - 1) / 2)
		}
	}
	fmt.Println("ansssss", ans)
	return ans
}
func func3() {
	// 输入：matrix = [[0,0,1],[1,1,1],[1,0,1]]
	// 输出：4
	// 解释：你可以按照上图方式重新排列矩阵的每一列。
	// 最大的全 1 子矩阵是上图中加粗的部分，面积为 4 。
	// 输入：matrix = [[1,0,1,0,1]]
	// 输出：3
	// 解释：你可以按照上图方式重新排列矩阵的每一列。
	// 最大的全 1 子矩阵是上图中加粗的部分，面积为 3 。
	largestSubmatrix([][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}})
	largestSubmatrix([][]int{{1, 0, 1, 0, 1}})
	largestSubmatrix([][]int{{0, 0}, {0, 0}})
	largestSubmatrix([][]int{{1, 1, 0}, {1, 0, 1}})
}
func largestSubmatrix(matrix [][]int) int {
	// 使用递增栈的方式
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// min := func(a, b int) int {
	// 	if a < b {0000
	// 		return a
	// 	}
	// 	return b
	// }
	m := len(matrix)
	n := len(matrix[0])
	// 将高度确认之后
	// 通过行遍历，找出最大的矩形
	for j := 0; j < n; j++ {
		height := 0
		for i := 0; i < m; i++ {
			if matrix[i][j] == 1 {
				height++
			} else {
				height = 0
			}
			matrix[i][j] = height
		}
	}
	// fmt.Println(matrix)
	ans := 0
	for i := 0; i < m; i++ {
		// 在这一行里面找到最高，或者最宽的矩形即可
		row := make([]int, n)
		copy(row, matrix[i])
		sort.Ints(row)
		cur := 0
		for j := 0; j < n; j++ {
			cur = max(cur, row[j]*(n-j))
		}
		fmt.Println("on row i ", i, row, cur)
		ans = max(cur, ans)
	}
	fmt.Println("ansssss", ans)
	return ans
}
func func4() {}
