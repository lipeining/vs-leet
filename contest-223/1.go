package main

import "fmt"

func main() {
	// func1()
	// func2()
	// fmt.Println("anssssssssssssssss")
	func3()
}
func func1() {
	// 输入：encoded = [1,2,3], first = 1
	// 输出：[1,0,2,1]
	// 解释：若 arr = [1,0,2,1] ，那么 first = 1 且 encoded = [1 XOR 0, 0 XOR 2, 2 XOR 1] = [1,2,3]
	// 示例 2：

	// 输入：encoded = [6,2,7,3], first = 4
	// 输出：[4,2,0,7,4]
	decode([]int{1, 2, 3}, 1)
	decode([]int{6, 2, 7, 3}, 4)
}
func decode(encoded []int, first int) []int {
	m := len(encoded)
	n := m + 1
	ans := make([]int, n)
	ans[0] = first
	for i := 1; i < n; i++ {
		ans[i] = encoded[i-1] ^ ans[i-1]
	}
	fmt.Println("ans", ans)
	return ans
}
func func2() {

}
func func4() {
	// 	输入：jobs = [3,2,3], k = 3
	// 输出：3
	// 解释：给每位工人分配一项工作，最大工作时间是 3 。
	// 示例 2：

	// 输入：jobs = [1,2,4,7,8], k = 2
	// 输出：11
	// 解释：按下述方式分配工作：
	// 1 号工人：1、2、8（工作时间 = 1 + 2 + 8 = 11）
	// 2 号工人：4、7（工作时间 = 4 + 7 = 11）
	// 最大工作时间是 11 。
	minimumTimeRequired([]int{3, 2, 3}, 3)
	minimumTimeRequired([]int{1, 2, 4, 7, 8}, 2)
}
func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	// 将 jobs 分为 k 个子集，每一个子集的和为该子集的值，
	// 求 k 个子集的最大值最小
	
	return n
}
func func3() {
	// 输入：source = [1,2,3,4], target = [2,1,4,5], allowedSwaps = [[0,1],[2,3]]
	// 输出：1
	// 解释：source 可以按下述方式转换：
	// - 交换下标 0 和 1 指向的元素：source = [2,1,3,4]
	// - 交换下标 2 和 3 指向的元素：source = [2,1,4,3]
	// source 和 target 间的汉明距离是 1 ，二者有 1 处元素不同，在下标 3 。
	// 示例 2：

	// 输入：source = [1,2,3,4], target = [1,3,2,4], allowedSwaps = []
	// 输出：2
	// 解释：不能对 source 执行交换操作。
	// source 和 target 间的汉明距离是 2 ，二者有 2 处元素不同，在下标 1 和下标 2 。
	// 示例 3：

	// 输入：source = [5,1,2,4,3], target = [1,5,4,2,3], allowedSwaps = [[0,4],[4,2],[1,3],[1,4]]
	// 输出：0
	minimumHammingDistance([]int{1, 2, 3, 4}, []int{2, 1, 4, 5}, [][]int{{0, 1}, {2, 3}})
	minimumHammingDistance([]int{1, 2, 3, 4}, []int{1, 3, 2, 4}, [][]int{})
	minimumHammingDistance([]int{5, 1, 2, 4, 3}, []int{1, 5, 4, 2, 3}, [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}})
}
func find(parent []int, i int) int {
	if parent[i] != i {
		parent[i] = find(parent, parent[i])
	}
	return parent[i]
}
func union(parent []int, rank []int, x int, y int) {
	x = find(parent, x)
	y = find(parent, y)
	if x != y {
		if rank[x] >= rank[y] {
			parent[y] = x
			rank[x] += rank[y]
		} else {
			parent[x] = y
			rank[y] += rank[x]
		}
	}
}
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	for _, swap := range allowedSwaps {
		union(parent, rank, swap[0], swap[1])
	}
	ans := 0
	// 对在同一个区间的进行不同数量的比较
	root := make(map[int]map[int]bool)
	for i := 0; i < n; i++ {
		r := find(parent, i)
		if _, ok := root[r]; !ok {
			root[r] = make(map[int]bool)
		}
		root[r][i] = true
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for _, set := range root {
		left := make(map[int]int)
		right := make(map[int]int)
		for index := range set {
			left[source[index]]++
			right[target[index]]++
		}
		same := 0
		for k, v := range left {
			if v2, ok := right[k]; ok {
				same += min(v, v2)
			}
		}
		// fmt.Println("set same number", set, same)
		ans += len(set) - same
	}
	// fmt.Println("ansssss", ans)
	return ans
}
