package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// func1()
	func2()
}
func func2() {
	getNumberOfBacklogOrders([][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}})
	getNumberOfBacklogOrders([][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}})
	getNumberOfBacklogOrders([][]int{{30, 27, 1}, {18, 9, 1}, {11, 4, 0}, {21, 11, 0}, {1, 1, 1}, {24, 20, 1}, {15, 13, 1}, {13, 3, 0}, {30, 11, 1}})
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

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
func getNumberOfBacklogOrders(orders [][]int) int {
	n := len(orders)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	buy := make(MaxHeap, 0)
	sell := make(MinHeap, 0)
	heap.Init(&buy)
	heap.Init(&sell)
	for i := 0; i < n; i++ {
		price, amount, t := orders[i][0], orders[i][1], orders[i][2]
		if t == 0 {
			// buy
			for amount > 0 {
				if sell.Len() > 0 {
					left := heap.Pop(&sell).(*Item)
					if left.priority > price {
						heap.Push(&sell, left)
						break
					}
					cost := min(left.value, amount)
					left.value -= cost
					amount -= cost
					if left.value > 0 {
						heap.Push(&sell, left)
					}
				} else {
					break
				}
			}
			if amount > 0 {
				heap.Push(&buy, &Item{
					value:    amount,
					priority: price,
					index:    i,
				})
			}
		} else {
			for amount > 0 {
				if buy.Len() > 0 {
					left := heap.Pop(&buy).(*Item)
					if left.priority < price {
						heap.Push(&buy, left)
						break
					}
					cost := min(left.value, amount)
					left.value -= cost
					amount -= cost
					if left.value > 0 {
						heap.Push(&buy, left)
					}
				} else {
					break
				}
			}
			if amount > 0 {
				heap.Push(&sell, &Item{
					value:    amount,
					priority: price,
					index:    i,
				})
			}
		}
		fmt.Println("on i", i)
		for j := 0; j < buy.Len(); j++ {
			fmt.Println("buy j", buy[j])
		}
		for j := 0; j < sell.Len(); j++ {
			fmt.Println("sell j", sell[j])
		}
	}
	ans := 0
	mod := int(1e9 + 7)
	for i := 0; i < buy.Len(); i++ {
		fmt.Println(buy[i])
		ans += buy[i].value
		ans %= mod
	}
	fmt.Println("after  buy", ans)
	for i := 0; i < sell.Len(); i++ {
		fmt.Println(sell[i])
		ans += sell[i].value
		ans %= mod
	}
	fmt.Println("after  sell", ans)
	fmt.Println("ansssssss", ans)
	return ans
}
func getNumberOfBacklogOrdersArr(orders [][]int) int {
	n := len(orders)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// 只找价格最低或者最高的，从这个开始
	for i := 1; i < n; i++ {
		rp, ra, rt := orders[i][0], orders[i][1], orders[i][2]
		for j := 0; j < i; j++ {
			lp, la, lt := orders[j][0], orders[j][1], orders[j][2]
			if rt == 0 {
				// right is buy
				if lt == 1 {
					if lp <= rp && ra > 0 && la > 0 {
						cost := min(ra, la)
						ra -= cost
						la -= cost
					}
				}
			} else {
				if lt == 0 {
					if lp >= rp && ra > 0 && la > 0 {
						cost := min(ra, la)
						ra -= cost
						la -= cost
					}
				}
			}
			orders[j][1] = la
		}
		orders[i][1] = ra
	}
	fmt.Println(orders)
	ans := 0
	mod := int(1e9 + 7)
	for i := 0; i < n; i++ {
		ans += orders[i][1]
		ans %= mod
	}
	fmt.Println("ansssss", ans)
	return ans
}
func func1() {
	maxAscendingSum([]int{10, 20, 30, 5, 10, 50})
	maxAscendingSum([]int{10, 20, 30, 40, 50})
	maxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12})
	maxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12})
	maxAscendingSum([]int{100, 10, 1})
	maxAscendingSum([]int{1})
}
func maxAscendingSum(nums []int) int {
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sum := nums[0]
	ans := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		ans = max(ans, sum)
	}
	fmt.Println("anssss", ans)
	return ans
}
