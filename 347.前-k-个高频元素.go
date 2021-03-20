/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */
package main

import (
	"sort"
)

// func main() {
// 	topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
// 	topKFrequent([]int{1, 1}, 1)
// 	topKFrequent([]int{3, 0, 1, 0}, 1)
// }

// @lc code=start

func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[nums[i]]++
	}
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	// fmt.Println(m)
	// fmt.Println("before sort", keys)
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	// fmt.Println("after sort", keys)
	return keys[:k]
}

// @lc code=end
