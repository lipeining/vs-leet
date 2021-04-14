package main

import "container/heap"

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	// return pq[i].priority > pq[j].priority
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func magicTower(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < 0 {
		return -1
	}
	now := 0
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	ans := 0
	for i, num := range nums {
		if num < 0 {
			heap.Push(&pq, &Item{value: num, priority: num, index: i})
		}
		now += num
		for now < 0 {
			item := heap.Pop(&pq).(*Item)
			now -= item.value
			ans++
		}
	}
	return ans
}
