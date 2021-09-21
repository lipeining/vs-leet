package main

import (
	"container/heap"
	"fmt"
)

func main() {
	minSwaps("][][")
	minSwaps("]]][[[")
	minSwaps("[]")
	minSwaps("")
}
func minSwaps(s string) int {
	n := len(s)
	ans := 0
	count := 0
	for i := 0; i < n; i++ {
		if s[i] == '[' {
			count++
		} else if count > 0 {
			count--
		} else {
			count++
			ans++
		}
	}
	fmt.Println("ans", ans)
	return ans
}

func isPrefixString(s string, words []string) bool {
	n := len(words)
	j := 0
	sl := len(s)
	for i := 0; i < n; i++ {
		word := words[i]
		end := j + len(word)
		if end > sl {
			return false
		}
		if s[j:end] != word {
			return false
		}
		if end == sl {
			return true
		}
		j = end
	}
	return false
}

type maxitem struct {
	t int
	i int
}
type maxheap []maxitem

func (h maxheap) Len() int            { return len(h) }
func (h maxheap) Less(i, j int) bool  { a, b := h[i], h[j]; return a.t > b.t }
func (h maxheap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxheap) Push(v interface{}) { *h = append(*h, v.(maxitem)) }
func (h *maxheap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *maxheap) push(v maxitem)     { heap.Push(h, v) }
func (h *maxheap) pop() maxitem       { return heap.Pop(h).(maxitem) }
func (h *maxheap) init()              { heap.Init(h) }

func minStoneSum(piles []int, k int) int {
	n := len(piles)
	pq := make(maxheap, n)
	for i := 0; i < n; i++ {
		pq[i] = maxitem{i: i, t: piles[i]}
	}
	pq.init()
	for k != 0 {
		if pq.Len() == 0 {
			break
		}
		top := pq.pop()
		top.t = top.t - top.t/2
		if top.t > 0 {
			pq.push(top)
		}
		k--
	}
	ans := 0
	for i := 0; i < pq.Len(); i++ {
		ans += pq[i].t
	}
	return ans
}
