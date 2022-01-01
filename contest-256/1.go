package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {}
func minSessions(tasks []int, sessionTime int) int {
	n := len(tasks)
	dp := make([][]int, n)
}
func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		return len(nums[i]) > len(nums[j]) || len(nums[i]) == len(nums[j]) && nums[i] > nums[j]
	})
	fmt.Println(nums)
	return nums[k-1]
}
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	if k == 1 {
		return nums[0]
	}
	ans := math.MaxInt32
	n := len(nums)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		j := i + k - 1
		if j >= n {
			break
		}
		ans = min(ans, nums[j]-nums[i])
	}
	return ans
}
