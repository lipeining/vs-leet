/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的K对数字
 */

// @lc code=start

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		tmp := make([][]int, 0)
		return tmp
	}
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for i := 0; i < len(nums1); i++ {
		// 一行一列的开始
		for j := 0; j < len(nums2); j++ {
			priority := nums1[i] + nums2[j]
			value := []int{nums1[i], nums2[j]}
			item := &Item{
				value:    value,
				priority: priority,
				index:    i*len(nums1) + j,
			}
			heap.Push(&pq, item)
			if pq.Len() > k {
				heap.Pop(&pq)
			}
		}
	}
	max := k
	if max > pq.Len() {
		max = pq.Len()
	}
	ans := make([][]int, max)
	i := max - 1
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		ans[i] = item.value
		i--
	}
	return ans
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    []int // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
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

// @lc code=end

