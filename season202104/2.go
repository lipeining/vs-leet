package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	func21()
	// func23()
}

type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A MaxHeap implements heap.Interface and holds Items.
type MaxHeap []*Item

func (pq MaxHeap) Len() int { return len(pq) }

func (pq MaxHeap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
	// return pq[i].priority < pq[j].priority
}

func (pq MaxHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MaxHeap) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *MaxHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
func storeWater(bucket []int, vat []int) int {
	ceil := func(want, per int) int {
		need := want / per
		if want%per != 0 {
			need++
		}
		return need
	}
	n := len(bucket)
	ans := 0
	pq := make(MaxHeap, 0)
	for i := 0; i < n; i++ {
		if vat[i] == 0 {
			continue
		} else {
			b := bucket[i]
			if bucket[i] == 0 {
				b++
				bucket[i]++
				ans++
			}
			pq = append(pq, &Item{
				value:    i,
				priority: ceil(vat[i], b),
				index:    i,
			})
		}
	}
	heap.Init(&pq)
	fmt.Println("base  anssssss", ans)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	cur := math.MaxInt32
	pull := 0
	for {
		top := heap.Pop(&pq).(*Item)
		fmt.Println(top, cur)
		// 有两种选择
		cur = min(cur, top.priority+pull)
		if top.priority <= 1 {
			break
		}
		pull++
		kindex := top.value
		b := bucket[kindex]
		np := ceil(vat[kindex], b+1)
		for np == top.priority {
			pull++
			b++
			bucket[kindex]++
			np = ceil(vat[kindex], b+1)
		}
		heap.Push(&pq, &Item{
			value:    kindex,
			priority: np,
			index:    kindex,
		})
	}
	fmt.Println("currrrrr", cur, "pull", pull)
	fmt.Println("pq")
	for i := 0; i < len(pq); i++ {
		fmt.Println(pq[i])
	}
	return ans + cur
}
func func21() {
	storeWater([]int{1, 3}, []int{6, 8})
	storeWater([]int{9, 0, 1}, []int{0, 2, 2})
	fmt.Println("2111111111111111111 anssssssssss")
}
func func22() {
	fmt.Println("22222222222222222222 anssssssssss")
}
func func23() {
	fmt.Println("2333333333333333333 anssssssssss")
}
func func24() {
	fmt.Println("24444444444444444444444 anssssssssss")
}
