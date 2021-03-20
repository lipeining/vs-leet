package main

import "fmt"

// func main() {
// 	searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
// 	searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
// }
func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	lowBound := func() int {
		l, r := 0, n-1
		for l < r {
			m := l + (r-l)/2
			fmt.Println("low", l, r, m, nums[m])
			if nums[m] >= target {
				r = m
			} else {
				l = m + 1
			}
		}
		fmt.Println("low end", l, r)
		if nums[l] == target {
			return l
		}
		return -1
	}
	upBound := func() int {
		l, r := 0, n-1
		for l < r {
			m := l + (r-l)/2
			fmt.Println("up", l, r, m, nums[m])
			if nums[m] <= target {
				l = m + 1
			} else {
				r = m
			}
		}
		fmt.Println("up end", l, r)
		if nums[l] == target {
			return l
		}
		if l > 0 && nums[l-1] == target {
			return l - 1
		}
		return -1
	}
	l := lowBound()
	if l == -1 {
		return []int{-1, -1}
	}
	r := upBound()
	fmt.Println("ans", l, r)
	return []int{l, r}
}
