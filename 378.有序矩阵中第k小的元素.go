/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第K小的元素
 */

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {

	n := len(matrix)
	if n == 0 {
		return 0
	}
	// Create a priority queue, put the items in it, and
    // establish the priority queue (heap) invariants.
    pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for i:=0;i<n;i++ {
		// 一行一列的开始
		for j:=0;j<n;j++ {
			priority := matrix[i][j]
			item := &Item{
				value:    priority,
				priority: priority,
				index:    i*n+j,
			}
			heap.Push(&pq, item)
			if pq.Len() > k {
				heap.Pop(&pq)
			}
		}
	}
	return heap.Pop(&pq).(*Item).value
}
// An Item is something we manage in a priority queue.
type Item struct {
    value    int // The value of the item; arbitrary.
    priority int    // The priority of the item in the queue.
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

