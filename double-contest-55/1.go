package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func maxAlternatingSum(nums []int) int64 {
	var odd int64 = -1
	var even int64 = 0
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	for _, num := range nums {
		odd, even = max(odd, even+int64(num)), max(even, odd-int64(num))
	}
	return odd
}
func main() {
	// removeOccurrences("daabcbaabcbc", "abc")
	// removeOccurrences("axxxxyyyyb", "xy")
	// removeOccurrences("hhvhvaahvahvhvaavhvaasshvahvaln", "hva")
	// // 	"hhvhvaahvahvhvaavhvaasshvahvaln"
	// // "hva"
	// canBeIncreasing([]int{2, 3, 1, 2})
}

type minitem struct{ t, i, m int }
type minheap []minitem

func (h minheap) Len() int { return len(h) }
func (h minheap) Less(i, j int) bool {
	a, b := h[i], h[j]
	if a.t != b.t {
		return a.t < b.t
	}
	return a.i < b.i
	// if a.i != b.i {
	// 	return a.i < b.i
	// }
	// return a.m < b.m
}
func (h minheap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minheap) Push(v interface{}) { *h = append(*h, v.(minitem)) }
func (h *minheap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *minheap) push(v minitem)     { heap.Push(h, v) }
func (h *minheap) pop() minitem       { return heap.Pop(h).(minitem) }
func (h *minheap) init()              { heap.Init(h) }

// t price i shop m movie
type MovieRentingSystem struct {
	// 已经借出去的电影也分1个堆
	// 需要标记已经借出去的电影
	// price movie shop
	// rent movie shop
	movieTable map[int][][]int
	rentsort   minheap
	price      map[int]map[int]int
	rent       map[int]map[int]bool
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	movieTable := make(map[int][][]int)
	price := make(map[int]map[int]int)
	for _, entry := range entries {
		shop, movie, p := entry[0], entry[1], entry[2]
		movieTable[movie] = append(movieTable[movie], []int{p, shop})
		if price[movie] == nil {
			price[movie] = make(map[int]int)
		}
		price[movie][shop] = p
	}
	for k := range movieTable {
		sort.Slice(movieTable[k], func(i, j int) bool {
			if movieTable[k][i][0] != movieTable[k][j][0] {
				return movieTable[k][i][0] < movieTable[k][j][0]
			}
			return movieTable[k][i][1] < movieTable[k][j][1]
		})
	}
	rentsort := make(minheap, 0)
	return MovieRentingSystem{
		movieTable: movieTable,
		rentsort:   rentsort,
		price:      price,
		rent:       make(map[int]map[int]bool),
	}
}

func (this *MovieRentingSystem) Search(movie int) []int {
	now := this.movieTable[movie]
	ans := make([]int, 0)
	if now == nil {
		return ans
	}
	j := 0
	r := this.rent[movie]
	size := len(now)
	for len(ans) < 5 && j < size {
		item := now[j]
		shop := item[1]
		j++
		// 如果没有被借
		if r != nil && r[shop] {
			continue
		}
		ans = append(ans, shop)
	}
	return ans
}

func (this *MovieRentingSystem) Rent(shop int, movie int) {
	if this.rent[movie] == nil {
		this.rent[movie] = make(map[int]bool)
	}
	this.rent[movie][shop] = true
	this.rentsort.push(minitem{t: this.price[movie][shop], i: shop, m: movie})
}

func (this *MovieRentingSystem) Drop(shop int, movie int) {
	delete(this.rent[movie], shop)
}

func (this *MovieRentingSystem) Report() [][]int {
	ans := make([][]int, 0)
	keep := make([]minitem, 0)
	for len(ans) < 5 && this.rentsort.Len() != 0 {
		item := this.rentsort.pop()
		movie, shop := item.m, item.i
		r := this.rent[movie]
		if r != nil && r[shop] {
			ans = append(ans, []int{shop, movie})
			keep = append(keep, item)
		}
	}
	for _, item := range keep {
		this.rentsort.push(item)
	}
	return ans
}

/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * obj := Constructor(n, entries);
 * param_1 := obj.Search(movie);
 * obj.Rent(shop,movie);
 * obj.Drop(shop,movie);
 * param_4 := obj.Report();
 */

func canBeIncreasing(nums []int) bool {
	n := len(nums)
	// 找第一个不合理的地方
	i := 0
	for i < n-1 {
		if nums[i] >= nums[i+1] {
			break
		} else {
			i++
		}
	}
	if i == n-1 {
		return true
	}
	// nums[i] >= nums[i+1]
	// 选择一个删除
	check := func(index int, big int) bool {
		for j := index; j < n; j++ {
			if nums[j] <= big {
				return false
			}
			big = nums[j]
		}
		return true
	}
	fmt.Println("on index", i)
	left := 0
	if i >= 1 {
		left = nums[i-1]
	}
	return check(i+1, left) || check(i+2, nums[i])
}
func removeOccurrences(s string, part string) string {
	m := len(part)
	for {
		n := len(s)
		i := 0
		fmt.Println(s)
		if n < m {
			break
		}
		if n == m && s != part {
			break
		}
		flag := false
		for i <= n-m {
			if s[i:i+m] == part {
				flag = true
				s = s[:i] + s[i+m:]
				break
			} else {
				i++
			}
		}
		if !flag {
			break
		}
	}
	return s
}
