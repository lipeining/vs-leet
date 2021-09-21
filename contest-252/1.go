package main

import (
	"container/heap"
)

func main() {
	numberOfWeeks([]int{5, 7, 5, 7, 9, 7})
}

type maxitem struct {
	t int64
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

func numberOfWeeks(milestones []int) int64 {
	return 0
}
