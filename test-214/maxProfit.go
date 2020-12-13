package main

import (
	"container/heap"
)

func main() {
	// 	输入：inventory = [2,5], orders = 4
	// 输出：14
	// 解释：卖 1 个第一种颜色的球（价值为 2 )，卖 3 个第二种颜色的球（价值为 5 + 4 + 3）。
	// 最大总和为 2 + 5 + 4 + 3 = 14 。
	// 示例 2：
	// maxProfit([]int{2, 5}, 4)
	// 输入：inventory = [3,5], orders = 6
	// 输出：19
	// 解释：卖 2 个第一种颜色的球（价值为 3 + 2），卖 4 个第二种颜色的球（价值为 5 + 4 + 3 + 2）。
	// 最大总和为 3 + 2 + 5 + 4 + 3 + 2 = 19 。
	// 示例 3：
	// maxProfit([]int{3, 5}, 6)
	// 输入：inventory = [2,8,4,10,6], orders = 20
	// 输出：110
	// 示例 4：
	// maxProfit([]int{2, 8, 4, 10, 6}, 20)
	// 输入：inventory = [1000000000], orders = 1000000000
	// 输出：21
	// 解释：卖 1000000000 次第一种颜色的球，总价值为 500000000500000000 。 500000000500000000 对 109 + 7 取余为 21 。
	maxProfit([]int{1000000000}, 1000000000)
}
func maxProfit(inventory []int, orders int) int {
	// sort.Ints(inventory)
	ans := 0
	toMod := int(1e9 + 7)
	// 贪心算法，从最大的开始计算吗？
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for _, invent := range inventory {
		value := invent
		priority := invent
		item := &Item{
			value:    value,
			priority: priority,
		}
		heap.Push(&pq, item)
	}
	for orders != 0 {
		item := heap.Pop(&pq).(*Item)
		var next *Item
		if pq.Len() != 0 {
			next = heap.Pop(&pq).(*Item)
			heap.Push(&pq, &Item{
				value:    next.value,
				priority: next.value,
			})
		}
		// fmt.Println(item, next)
		use := min(item.value, orders)
		if next != nil {
			if item.value != next.value {
				use = item.value - next.value
			} else {
				use = 1
			}
		}
		ans = ans + count(item.value, use)
		ans %= toMod
		if item.value > use {
			heap.Push(&pq, &Item{
				value:    item.value - use,
				priority: item.value - use,
			})
		}
		orders -= use
	}
	// fmt.Println("ans", ans)
	// fmt.Println("-----------")
	return ans
}
func count(start, use int) int {
	toMod := int(1e9 + 7)
	ans := 0
	// fmt.Println("count, use", start, use)
	for i := 0; i < use; i++ {
		ans += start
		ans %= toMod
		start--
	}
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
	// return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
