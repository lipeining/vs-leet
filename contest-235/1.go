package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// minAbsoluteSumDiff([]int{1, 7, 5}, []int{2, 3, 5})
	// minAbsoluteSumDiff([]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10})
	// minAbsoluteSumDiff([]int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4})
	countDifferentSubsequenceGCDs([]int{6, 10, 3})
	countDifferentSubsequenceGCDs([]int{5, 15, 40, 5, 6})
}

func countDifferentSubsequenceGCDs(nums []int) int {
	// 可能是去重，然后通过排序，
}
func countDifferentSubsequenceGCDsDFS(nums []int) int {
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}
	sort.Ints(nums)
	c := make(map[int]bool)
	var dfs func(path []int, g int, index int)
	dfs = func(path []int, g int, index int) {
		// 计算当前的 path 的 最大 gcd
		if len(path) != 0 {
			fmt.Println("for path, g ", path, g)
			c[g] = true
		}
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			if g == 0 {
				dfs(path, nums[i], i+1)
			} else {
				ng := gcd(g, nums[i])
				dfs(path, ng, i+1)
			}
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	dfs(path, 0, 0)
	ans := len(c)
	fmt.Println(c)
	fmt.Println("anssssss", ans)
	return ans
}
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	mod := int(1e9 + 7)
	ans := 0
	n := len(nums1)
	for i := 0; i < n; i++ {
		ans += abs(nums1[i] - nums2[i])
		// ans %= mod
	}
	if ans == 0 {
		return ans
	}
	s1 := make([]int, n)
	copy(s1, nums1)
	sort.Ints(s1)
	// fmt.Println(s1)
	bin := func(target int) int {
		l := 0
		r := n - 1
		for l < r-1 {
			mid := l + (r-l)/2
			if s1[mid] == target {
				return target
			} else if s1[mid] < target {
				l = mid
			} else {
				r = mid
			}
		}
		// fmt.Println("for target", target, s1[l], s1[r])
		if abs(s1[l]-target) < abs(s1[r]-target) {
			return s1[l]
		}
		return s1[r]
	}
	mm := ans
	// fmt.Println("base ans", ans)
	for i := 0; i < n; i++ {
		a := abs(nums1[i] - nums2[i])
		if a == 0 {
			continue
		}
		// 找到一个最接近 nums2[i] 的值进行替换，计算之后的 nums
		target := nums2[i]
		want := bin(target)
		nans := ans - a + abs(want-target)
		// fmt.Println("for i", i, "target", target, "want", want, "nans", nans)
		if nans < mm {
			mm = nans
		}
	}
	if ans > mm {
		ans = mm
	}
	fmt.Println("anssssssss", ans)
	ans %= mod
	return ans
}
func findingUsersActiveMinutes(logs [][]int, k int) []int {
	ans := make([]int, k)
	c := make(map[int]map[int]bool)
	// 用户在 time 时的活跃
	for _, log := range logs {
		id, time := log[0], log[1]
		if c[id] == nil {
			c[id] = make(map[int]bool)
		}
		c[id][time] = true
	}
	fmt.Println(c)
	for _, tv := range c {
		num := len(tv)
		ans[num-1]++
	}
	return ans
}
func truncateSentence(s string, k int) string {
	arr := strings.Split(s, " ")
	arr = arr[:k]
	return strings.Join(arr, " ")
}
