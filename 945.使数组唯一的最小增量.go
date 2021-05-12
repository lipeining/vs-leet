/*
 * @lc app=leetcode.cn id=945 lang=golang
 *
 * [945] 使数组唯一的最小增量
 */
package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	minIncrementForUnique([]int{3, 2, 1, 2, 1, 7})
// }

// @lc code=start
func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	ans := 0
	n := len(A)
	fmt.Println(A)
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for i := 1; i < n; i++ {
		if A[i] <= A[i-1] {
			diff := A[i] - A[i-1]
			want := 1 + abs(diff)
			ans += want
			A[i] += want
		} else {
			// empty
		}
	}
	fmt.Println(A, ans)
	return ans
}

// @lc code=end
