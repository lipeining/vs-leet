package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// eliminateMaximum([]int{1, 3, 4}, []int{1, 1, 1})
	// eliminateMaximum([]int{3, 2, 4}, []int{5, 3, 2})
	// eliminateMaximum([]int{1, 1, 2, 3}, []int{1, 1, 1, 1})
	countGoodNumbers(1)
	countGoodNumbers(4)
	countGoodNumbers(50)
}
func countGoodNumbers(n int64) int {
	toMod := int(1e9 + 7)
	memo := make(map[int64]int)
	var dfs func(num int64) int
	dfs = func(num int64) int {
		if num == 1 {
			return 5
		}
		if num == 2 {
			return 20
		}
		if memo[num] != 0 {
			return memo[num]
		}
		h := num / 2
		half := dfs(h)
		now := half * half % toMod
		if num%2 == 1 {
			now = now * 5 % toMod
		}
		memo[num] = now
		return now
	}
	ans := dfs(n)
	fmt.Println("ansss", ans)
	return ans
}

type minitem struct{ t, i int }
type minheap []minitem

func (h minheap) Len() int            { return len(h) }
func (h minheap) Less(i, j int) bool  { a, b := h[i], h[j]; return a.t < b.t }
func (h minheap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minheap) Push(v interface{}) { *h = append(*h, v.(minitem)) }
func (h *minheap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *minheap) push(v minitem)     { heap.Push(h, v) }
func (h *minheap) pop() minitem       { return heap.Pop(h).(minitem) }
func (h *minheap) init()              { heap.Init(h) }
func eliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	ceil := func(a, b int) int {
		t := a / b
		if a%b != 0 {
			t++
		}
		return t
	}
	pq := make(minheap, n)
	// t 表示时间
	for i := 0; i < n; i++ {
		d, s := dist[i], speed[i]
		t := ceil(d, s)
		pq[i] = minitem{t: t, i: i}
	}
	pq.init()
	t := 0
	ans := 0
	for pq.Len() != 0 {
		top := pq.pop()
		if top.t <= t {
			fmt.Println("fail at", top)
			break
		}
		fmt.Println("choose, i", top)
		ans++
		t++
	}
	fmt.Println("anssss", ans)
	return ans
}
func buildArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}
