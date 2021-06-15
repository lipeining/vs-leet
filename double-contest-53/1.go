package main

import (
	"fmt"
	"sort"
)

func main() {
	getBiggestThree([][]int{
		{3, 4, 5, 1, 3},
		{3, 3, 4, 2, 3},
		{20, 30, 200, 40, 10},
		{1, 5, 5, 4, 1},
		{4, 3, 2, 2, 5},
	})
}
func getBiggestThree(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	vet := func(i, j int) int {
		size := 1
		ret := grid[i][j]
		for {
			now := 0
			fmt.Println("size", size, i, j, ret)
			fmt.Print("from ", i, "-", j, "=>")
			movei, movej := i, j
			// up to left
			for x := 0; x < size; x++ {
				movei = movei + x + 1
				movej = movej - x - 1
				if movei < 0 || movei >= m || movej < 0 || movej >= n {
					return ret
				}
				fmt.Print(movei, "-", movej, "=>")
				now += grid[movei][movej]
			}
			// left to buttom
			for x := 0; x < size; x++ {
				movei = movei + x + 1
				movej = movej + x + 1
				if movei < 0 || movei >= m || movej < 0 || movej >= n {
					return ret
				}
				fmt.Print(movei, "-", movej, "=>")
				now += grid[movei][movej]
			}
			// buttom to right
			for x := 0; x < size; x++ {
				movei = movei - x - 1
				movej = movej + x + 1
				if movei < 0 || movei >= m || movej < 0 || movej >= n {
					return ret
				}
				fmt.Print(movei, "-", movej, "=>")
				now += grid[movei][movej]
			}
			// right to up
			for x := 0; x < size; x++ {
				movei = movei - x - 1
				movej = movej - x - 1
				if movei < 0 || movei >= m || movej < 0 || movej >= n {
					return ret
				}
				fmt.Print(i, "-", j, "=>")
				now += grid[movei][movej]
			}
			fmt.Println("size loop")
			ret = now
			size++
		}
	}
	keys := make(map[int]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := vet(i, j)
			fmt.Println("i, j", i, j, v)
			keys[v] = true
		}
	}
	ans := make([]int, 0)
	for k := range keys {
		ans = append(ans, k)
	}
	sort.Ints(ans)
	// fmt.Println("ansssss", ans)
	arr := make([]int, 0)
	for i := len(ans) - 1; i >= 0 && len(arr) < 3; i-- {
		arr = append(arr, ans[i])
	}
	// fmt.Println("arrrr", arr)
	return arr
}
func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := 0
	for i := 0; i < n/2; i++ {
		sum := nums[i] + nums[n-i-1]
		ans = max(ans, sum)
	}
	return ans
}
func countGoodSubstrings(s string) int {
	window := 3
	cnt := make(map[byte]int)
	n := len(s)
	l := 0
	r := 0
	total := 0
	ans := 0
	for r < window {
		if cnt[s[r]] == 0 {
			total++
		}
		cnt[s[r]]++
		r++
	}
	for r < n {
		if total >= 3 {
			ans++
		}
		if cnt[s[r]] == 0 {
			total++
		}
		cnt[s[r]]++
		r++
		if cnt[s[l]] == 1 {
			total--
		}
		cnt[s[l]]--
		l++
	}

	return ans
}
