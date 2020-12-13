package main

import "fmt"

func bestTeamScore(scores []int, ages []int) int {
	// 回溯算法
	n := len(scores)
	ans := 0
	var dfs func(index int, sum int, path []int)
	dfs = func(index int, sum int, path []int) {
		// 如何剪枝叶
		fmt.Println(index, sum, path)
		ans = max(ans, sum)
		if index == n {
			return
		}
		for i := index; i < n; i++ {
			// 是否可以取当前这个 i 用户
			// path 加上 i 之后，队内没有矛盾就可以加上 i
			flag := false
			for _, j := range path {
				// path 本来就是无矛盾的
				if ages[i] < ages[j] && scores[i] > scores[j] {
					flag = true
					break
				}
				if ages[j] < ages[i] && scores[j] > scores[i] {
					flag = true
				}
			}
			if flag {
				continue
			}
			path = append(path, i)
			dfs(i+1, sum+scores[i], path)
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	dfs(0, 0, path)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func test5545() {
	// 示例 1：

	// 输入：scores = [1,3,5,10,15], ages = [1,2,3,4,5]
	// 输出：34
	// 解释：你可以选中所有球员。

	fmt.Println(bestTeamScore([]int{1, 3, 5, 10, 15}, []int{1, 2, 3, 4, 5}))
	fmt.Println("------------------")
	// 示例 2：

	// 输入：scores = [4,5,6,5], ages = [2,1,2,1]
	// 输出：16
	// 解释：最佳的选择是后 3 名球员。注意，你可以选中多个同龄球员。
	fmt.Println(bestTeamScore([]int{4, 5, 6, 5}, []int{2, 1, 2, 1}))
	fmt.Println("------------------")
	// 示例 3：

	// 输入：scores = [1,2,3,5], ages = [8,9,10,1]
	// 输出：6
	// 解释：最佳的选择是前 3 名球员。
	fmt.Println(bestTeamScore([]int{1, 2, 3, 5}, []int{8, 9, 10, 1}))
	fmt.Println("------------------")

	// 	输入：
	// [9,2,8,8,2]
	// [4,1,3,3,5]
	// 输出：
	// 29
	// 预期：
	// 27
	fmt.Println(bestTeamScore([]int{9, 2, 8, 8, 2}, []int{4, 1, 3, 3, 5}))
	fmt.Println("------------------")

	// 输入：
	// [1,1,1,1,1,1,1,1,1,1]
	// [811,364,124,873,790,656,581,446,885,134]
	// 输出：
	// 1
	// 预期：
	// 10
	// fmt.Println(bestTeamScore([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []int{811, 364, 124, 873, 790, 656, 581, 446, 885, 134}))
	// fmt.Println("------------------")

	// 	[591,836,413,49,344,383,678,738,188,201,150,101,849,482,396,197,819,281,520,99,509,448,408,776,61,884,728,465,428,108,307,839,249,8,72,798,824,981,188,528,853,319,99,599,321,409,596,636,983,952,72,460,117,760,408,207,868,62,175,497,224,726,105,809,799,987,65,276,148,751,956,779]
	// [49,100,36,98,20,11,28,99,41,39,9,35,86,4,15,14,41,93,98,49,9,36,37,51,75,17,47,17,32,37,67,7,92,82,2,30,67,85,62,19,81,38,14,98,34,64,82,70,86,56,78,71,49,30,45,74,58,39,21,60,75,29,57,47,35,62,28,92,52,28,7,97]
}
