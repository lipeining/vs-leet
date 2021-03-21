package template

import "container/heap"

// pair 可以设置任意需要的字段
type pair struct{ t, i int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.t < b.t || a.t == b.t && a.i < b.i }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v pair)        { heap.Push(h, v) }
func (h *hp) pop() pair          { return heap.Pop(h).(pair) }

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

// // An Item is something we manage in a priority queue.
// type Item struct {
// 	value    int // The value of the item; arbitrary.
// 	priority int // The priority of the item in the queue.
// 	// The index is needed by update and is maintained by the heap.Interface methods.
// 	index int // The index of the item in the heap.
// }

// A MinHeap implements heap.Interface and holds Items.
type MinHeap []*Item

func (pq MinHeap) Len() int { return len(pq) }

func (pq MinHeap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	// return pq[i].priority > pq[j].priority
	return pq[i].priority < pq[j].priority
}

func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MinHeap) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *MinHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// // An Item is something we manage in a priority queue.
// type Item struct {
// 	value    int // The value of the item; arbitrary.
// 	priority int // The priority of the item in the queue.
// 	// The index is needed by update and is maintained by the heap.Interface methods.
// 	index int // The index of the item in the heap.
// }

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