package main

import (
	"fmt"
)

func main() {
	// subsetXORSum([]int{1, 3})
	// subsetXORSum([]int{5, 1, 6})
	// minSwaps("111000")
	// minSwaps("1110010")
	// minSwaps("1110001")
	// minSwaps("010")
	// minSwaps("1110")
}

func rearrangeSticks(n int, k int) int {
	// nums := make([]int, n+1)
	// dp 算法，在 index 下标处，能够看到的数量为k的方案数
	// dp[index][k] 那么 dp[index-1][k-1]
	// 也就是说，index 那里是倒数第 K 长的木棍
	// 而 index 前面的只能有
	// 答案应该是 dp[0-n][k] 的总和
	
	
	
	return 0
}

type FindSumPairs struct {
	nums1Map map[int]int
	nums2Map map[int]int
	nums2    []int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	nums1Map := make(map[int]int)
	for _, num := range nums1 {
		nums1Map[num]++
	}
	nums2Map := make(map[int]int)
	for _, num := range nums2 {
		nums2Map[num]++
	}
	return FindSumPairs{nums1Map: nums1Map, nums2Map: nums2Map, nums2: nums2}
}

func (this *FindSumPairs) Add(index int, val int) {
	origin := this.nums2[index]
	this.nums2Map[origin]--
	now := this.nums2[index] + val
	this.nums2Map[now]++
	this.nums2[index] = now
}

func (this *FindSumPairs) Count(tot int) int {
	cnt := 0
	for left, lc := range this.nums1Map {
		target := tot - left
		if rc, ok := this.nums2Map[target]; ok {
			// 组合数量
			cnt += lc * rc
		}
	}
	return cnt
}

func minSwaps(s string) int {
	zero := 0
	one := 0
	for _, char := range s {
		if char == '0' {
			zero++
		} else {
			one++
		}
	}
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
	if abs(zero-one) >= 2 {
		return -1
	}
	// odd := len(s)%2 == 1
	n := len(s)
	cacl := func(want byte) int {
		cnt := 0
		for i := 0; i < n; i += 2 {
			if s[i] != want {
				cnt++
			}
		}
		return cnt
	}
	ans := 0
	if zero > one {
		// 需要从 0 开始
		ans = cacl('0')
	} else if one > zero {
		ans = cacl('1')
	} else {
		ans = min(cacl('0'), cacl('1'))
	}
	fmt.Println("anssssss", ans)
	return ans
}
func subsetXORSum(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := 0
	// var dfs func(index int, xor int, path []int)
	// dfs = func(index int, xor int, path []int) {
	// 	cur := xor ^ 0
	// 	ans += cur
	// 	fmt.Println("cur", path, cur)
	// 	for i := index; i < n; i++ {
	// 		path = append(path, nums[i])
	// 		dfs(i+1, xor^nums[i], path)
	// 		path = path[:len(path)-1]
	// 	}
	// }
	// path := make([]int, 0)
	// dfs(0, 0, path)
	var dfs func(index int, xor int)
	dfs = func(index int, xor int) {
		cur := xor ^ 0
		ans += cur
		fmt.Println("cur", cur)
		for i := index; i < n; i++ {
			dfs(i+1, xor^nums[i])
		}
	}
	dfs(0, 0)
	fmt.Println("anssss", ans)
	return ans
}
