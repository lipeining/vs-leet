/*
 * @lc app=leetcode.cn id=1482 lang=golang
 *
 * [1482] 制作 m 束花所需的最少天数
 */
package main

import "fmt"

// import "fmt"

// func main() {
// 	minDays([]int{1, 10, 3, 10, 2}, 3, 1)
// }

// @lc code=start
func minDays1582(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if n < m*k {
		return -1
	}
	// max := func(a, b int) int {
	// 	if a > b {
	// 		return a
	// 	}
	// 	return b
	// }
	check := func(day int) bool {
		cnt := 0
		i := 0
		j := 0
		for i < n && cnt < m {
			j = i
			cur := 0
			for j < n {
				if bloomDay[j] <= day {
					j++
					cur++
				} else {
					break
				}
			}
			cnt += cur / k
			i = j + 1
		}
		return cnt >= m
	}
	l := 1
	r := int(1e9)
	for l < r {
		mid := l + (r-l)/2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Println("anssssss", r)
	if check(r) {
		return r
	}
	return -1
}

// @lc code=end
