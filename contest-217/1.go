package main

import (
	"fmt"
	"math"
)

func main() {
	// 	输入：nums = [3,5,2,6], k = 2
	// 	输出：[2,6]
	// 	解释：在所有可能的子序列集合 {[3,5], [3,2], [3,6], [5,2], [5,6], [2,6]} 中，[2,6] 最具竞争力。
	// 	示例 2：
	// mostCompetitive([]int{3, 5, 2, 6}, 2)
	// mostCompetitive([]int{2, 4, 3, 3, 5, 4, 9, 6}, 4)
	// mostCompetitive([]int{2, 1, 3, 1, 1, 1, 1, 6}, 4)
	// mostCompetitive([]int{2, 4, 3, 3, 5, 1, 1, 1}, 4)
	mostCompetitive([]int{6, 1, 10, 2, 6, 7, 9, 9, 7, 2, 6, 9, 9, 7, 8, 0, 0, 9, 6, 4, 1, 2, 6, 5, 2, 10, 3, 8, 8, 8, 9, 2, 9, 10, 6, 8, 1, 9, 5, 8, 10, 5, 4, 2, 5, 3, 1, 6, 7, 3, 3, 9, 5, 6, 10, 7, 0, 3, 1, 0, 5, 3}, 25)
	// 	输入：nums = [2,4,3,3,5,4,9,6], k = 4
	// 	输出：[2,3,3,4]
	// 	[6,1,10,2,6,7,9,9,7,2,6,9,9,7,8,0,0,9,6,4,1,2,6,5,2,10,3,8,8,8,9,2,9,10,6,8,1,9,5,8,10,5,4,2,5,3,1,6,7,3,3,9,5,6,10,7,0,3,1,0,5,3]
	// 25
	// 输出：
	// [0,0,1,1,5,5,2,5,3,1,6,7,3,3,9,5,6,10,7,0,3,1,0,5,3]
	// 预期：
	// [0,0,1,1,5,4,2,5,3,1,6,7,3,3,9,5,6,10,7,0,3,1,0,5,3]
}
func mostCompetitive(nums []int, k int) []int {
	ans := make([]int, k)
	n := len(nums)
	prevNode := -1
	// lastVal := math.MaxInt32
	// 从 0 到 n-k+i
	for i := 0; i < k; i++ {
		end := n - k + i
		findNode := -1
		prev := math.MaxInt32
		fmt.Println("start prevNode+1, end", prevNode+1, end)
		for j := prevNode + 1; j <= end; j++ {
			// if nums[j] == lastVal {
			// 	prev = nums[j]
			// 	findNode = j
			// 	break
			// }
			if nums[j] < prev {
				prev = nums[j]
				findNode = j
			}
		}
		fmt.Println("end prev, prevNode", prev, prevNode)
		prevNode = findNode
		ans[i] = prev
		// lastVal = prev
	}
	fmt.Println("ans-------------------", ans)
	return ans
}
func mostCompetitiveBack(nums []int, k int) []int {
	ans := make([]int, k)
	n := len(nums)
	prevNode := -1
	for i := 0; i < k; i++ {
		j := n - k + i
		prev := math.MaxInt32
		findNode := -1
		fmt.Println("start j", j, "prevNode", prevNode)
		for j > prevNode {
			if nums[j] <= prev {
				prev = nums[j]
				findNode = j
			}
			j--
		}
		fmt.Println("end j", j, prev, prevNode)
		prevNode = findNode
		ans[i] = prev
	}
	fmt.Println("ans-------------------", ans)
	return ans
}
