package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// splitString("1234")
	// splitString("050043")
	// splitString("9080701")
	// splitString("10009998")
	minInterval([][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5})
}
func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		dis := (intervals[i][1] - intervals[i][0]) - (intervals[j][1] - intervals[j][0])
		if dis == 0 {
			return intervals[i][1] < intervals[j][1]
		}
		return dis < 0
	})
	// 按照长度，右边界 排序
	fmt.Println(intervals)
	n := len(queries)
	ans := make([]int, n)
	binary := func(target int) int {
		l := 0
		r := len(intervals) - 1
		for l < r {
			fmt.Println("l, r", l, "-", r)
			mid := l + (r-l)/2
			interval := intervals[mid]
			left := interval[0]
			right := interval[1]
			if target < left {
				r = mid
			} else if target > right {
				l = mid + 1
			} else {
				// 符合条件，半边推进
				r = mid
			}
		}
		// 此时 l == r
		interval := intervals[l]
		left := interval[0]
		right := interval[1]
		if target >= left && target <= right {
			return right - left + 1
		}
		return -1
	}
	for i := 0; i < n; i++ {
		ans[i] = binary(queries[i])
		fmt.Println(" ----------- ")
	}
	fmt.Println("anssss", ans)
	return ans
}
func splitString(s string) bool {
	n := len(s)
	var dfs func(index int, prev int) bool
	dfs = func(index int, prev int) bool {
		if index >= n {
			return true
		}
		// 得到 cur 判断是否是第一个数字
		if prev == -1 {
			for j := index + 1; j < n; j++ {
				cur, _ := strconv.Atoi(s[index:j])
				if dfs(j, cur) {
					fmt.Println("-1", cur)
					return true
				}
			}
		} else {
			for j := index + 1; j <= n; j++ {
				cur, _ := strconv.Atoi(s[index:j])
				if prev-cur == 1 {
					if dfs(j, cur) {
						fmt.Println("prev", prev, "cur", cur)
						return true
					}
				}
			}
		}
		fmt.Println("ans false", prev)
		return false
	}
	ans := dfs(0, -1)
	fmt.Println("anssssss", ans)
	return ans
}
func getMinDistance(nums []int, target int, start int) int {
	n := len(nums)
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	ans := n
	for i := 0; i < n; i++ {
		if nums[i] == target {
			ans = min(ans, abs(i-start))
		}
	}
	return ans
}
