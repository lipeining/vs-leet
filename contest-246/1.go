package main

import "strconv"

func main() {

}
func minDifference(nums []int, queries [][]int) []int {
	C := 100
	n := len(nums)
	pre := make([][]int, 0)
	pre = append(pre, make([]int, C+1))
	for i := 0; i < n; i++ {
		now := make([]int, C+1)
		copy(now, pre[i])
		pre = append(pre, now)
		pre[i+1][nums[i]] = pre[i][nums[i]] + 1
	}
	m := len(queries)
	ans := make([]int, m)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for k := 0; k < m; k++ {
		last := 0
		best := C + 1
		l := queries[k][0]
		r := queries[k][1]
		for i := 1; i <= C; i++ {
			if pre[r+1][i] != pre[l][i] {
				if last != 0 {
					best = min(best, i-last)
				}
				last = i
			}
		}
		if best == C+1 {
			best = -1
		}
		ans[k] = best
	}
	return ans
}
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	// 暴力来一次
	m := len(grid1)
	n := len(grid1[0])
	g1 := 2
	g2 := 2
	t1 := make(map[int]map[int]bool)
	t2 := make(map[int]map[int]bool)
	id1 := make(map[int]int)
	id2 := make(map[int]int)
	toID := func(i, j int) int {
		return i*n + j
	}
	var dfs1 func(i, j int)
	dfs1 = func(i, j int) {
		if i < 0 || i == m || j < 0 || j == n {
			return
		}
		if grid1[i][j] != 1 {
			return
		}
		id := toID(i, j)
		if t1[g1] == nil {
			t1[g1] = make(map[int]bool)
		}
		t1[g1][id] = true
		id1[id] = g1
		grid1[i][j] = g1
		dfs1(i-1, j)
		dfs1(i, j-1)
		dfs1(i+1, j)
		dfs1(i, j+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == 1 {
				dfs1(i, j)
				g1++
			}
		}
	}
	var dfs2 func(i, j int)
	dfs2 = func(i, j int) {
		if i < 0 || i == m || j < 0 || j == n {
			return
		}
		if grid2[i][j] != 1 {
			return
		}
		id := toID(i, j)
		if t2[g2] == nil {
			t2[g2] = make(map[int]bool)
		}
		t2[g2][id] = true
		id2[id] = g2
		grid2[i][j] = g2
		dfs2(i-1, j)
		dfs2(i, j-1)
		dfs2(i+1, j)
		dfs2(i, j+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				dfs2(i, j)
				g2++
			}
		}
	}
	// 此时检查是否有覆盖的情况
	ans := 0
	iter2 := 2
	include := func(m1, m2 map[int]bool) bool {
		for k := range m2 {
			if _, ok := m1[k]; !ok {
				return false
			}
		}
		return true
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == iter2 {
				m2 := t2[iter2]
				iter1 := id1[toID(i, j)]
				m1 := t1[iter1]
				if include(m1, m2) {
					ans++
				}
				iter2++
			}
		}
	}
	return ans
}
func numberOfRounds(startTime string, finishTime string) int {
	toMin := func(time string) int {
		hour, _ := strconv.Atoi(time[0:2])
		minute, _ := strconv.Atoi(time[3:5])
		return hour*60 + minute
	}
	zone := 15
	ceil := func(a int) int {
		ans := a / zone
		if a%zone != 0 {
			ans++
		}
		return ans
	}
	oneDay := 60 * 24
	start := toMin(startTime)
	finish := toMin(finishTime)
	if start > finish {
		finish += oneDay
	}
	s := ceil(start)
	f := finish / zone
	ans := f - s
	return ans
}
func largestOddNumber(num string) string {
	// 以 i 结尾奇数
	n := len(num)
	ans := ""
	for i := n - 1; i >= 0; i-- {
		c := int(num[i] - '0')
		if c%2 == 1 {
			ans = num[:i]
			break
		}
	}
	return ans
}
