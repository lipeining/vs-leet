/*
 * @lc app=leetcode.cn id=973 lang=golang
 *
 * [973] 最接近原点的 K 个点
 */

// @lc code=start
func kClosest(points [][]int, K int) [][]int {
    // Create a priority queue, put the items in it, and
    // establish the priority queue (heap) invariants.
    pq := make(PriorityQueue, 0)
	heap.Init(&pq)
    for index,point := range points {
		priority := point[0]*point[0] + point[1]*point[1]
        item := &Item{
            value:    point,
            priority: priority,
            index:    index,
		}
		heap.Push(&pq, item)
		if pq.Len() > K {
			heap.Pop(&pq)
		}
    }
	ans := make([][]int, 0)
	for pq.Len()>0 {
		item := heap.Pop(&pq).(*Item)
		ans = append(ans, item.value)
	}
	return ans
}
// An Item is something we manage in a priority queue.
type Item struct {
    value    []int // The value of the item; arbitrary.
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

