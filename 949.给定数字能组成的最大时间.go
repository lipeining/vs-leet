/*
 * @lc app=leetcode.cn id=949 lang=golang
 *
 * [949] 给定数字能组成的最大时间
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// largestTimeFromDigits([]int{1, 2, 3, 4})
	// largestTimeFromDigits([]int{5, 5, 5, 5})
	// largestTimeFromDigits([]int{0, 0, 0, 0})
	// largestTimeFromDigits([]int{0, 0, 1, 0})
	largestTimeFromDigits([]int{1, 1, 1, 0})
}

// @lc code=start
func largestTimeFromDigits(arr []int) string {
	n := len(arr)
	used := make([]bool, n)
	ans := make([]int, 0)
	format := func(nums []int) bool {
		if nums[0] >= 3 {
			return false
		}
		if nums[0] == 2 {
			if nums[1] >= 4 {
				return false
			}
		}
		if nums[2] >= 6 {
			return false
		}
		return true
	}
	big := func(left, right []int) bool {
		if !format(left) {
			return false
		}
		if len(right) == 0 {
			return true
		}
		for i := 0; i < n; i++ {
			if left[i] < right[i] {
				return false
			} else if left[i] > right[i] {
				return true
			} else {
				continue
			}
		}
		return true
	}
	var dfs func(s []int)
	dfs = func(s []int) {
		if len(s) == n {
			if big(s, ans) {
				// fmt.Println(s, ans)
				t := make([]int, len(s))
				copy(t, s)
				ans = t
			}
			return
		}
		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				s = append(s, arr[i])
				dfs(s)
				used[i] = false
				s = s[:len(s)-1]
			}
		}
	}
	s := make([]int, 0)
	dfs(s)
	// fmt.Println("big", big([]int{1, 1, 1, 0}, []int{1, 1, 0, 1}))
	// fmt.Println("big", big([]int{1, 1, 0, 1}, []int{1, 1, 1, 0}))
	fmt.Println("anssss", ans)
	if len(ans) == 0 {
		return ""
	}
	ret := ""
	for i := 0; i < n; i++ {
		ret += strconv.Itoa(ans[i])
		if i == 1 {
			ret += ":"
		}
	}
	return ret
}

// @lc code=end
