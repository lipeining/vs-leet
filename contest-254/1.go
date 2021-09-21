package main

import (
	"sort"
	"strings"
)

func main() {

}
func numOfStrings(patterns []string, word string) int {
	ans := 0
	n := len(patterns)
	for i := 0; i < n; i++ {
		index := strings.Index(word, patterns[i])
		if index != -1 {
			ans++
		}
	}
	return ans
}

// 12345 1 2 4 3 5
func rearrangeArray(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([]int, 0)
	m := (n + 1) / 2
	for i := 0; i < m; i++ {
		ans = append(ans, nums[i])
		if i+m < n {
			ans = append(ans, nums[i+m])
		}
	}
	return nums
}
