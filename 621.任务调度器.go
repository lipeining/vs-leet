import "container/heap"

/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
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

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
func leastInterval(tasks []byte, n int) int {
	counter := make([]int, 26)
	for _, task := range tasks {
		counter[task-'A']++
	}
	pq := make(PriorityQueue, 0)
	for index, priority := range counter {
		if priority > 0 {
			// Insert a new item and then modify its priority.
			item := &Item{
				value:    index,
				priority: priority,
			}
			heap.Push(&pq, item)
		}
	}
	heap.Init(&pq)
	time := 0
	for pq.Len() != 0 {
		i := 0
		temp := make([]*Item, 0)
		for i <= n {
			if pq.Len() != 0 {
				top := heap.Pop(&pq).(*Item)
				if top.priority > 1 {
					temp = append(temp, top)
				}
			}
			time++
			if pq.Len() == 0 && len(temp) == 0 {
				break
			}
			i++
		}
		for _, oldItem := range temp {
			item := &Item{
				value:    oldItem.index,
				priority: oldItem.priority - 1,
			}
			heap.Push(&pq, item)
		}
	}
	return time
}

// @lc code=end

// public int leastInterval(char[] tasks, int n) {
// 	int[] map = new int[26];
// 	for (char c: tasks)
// 		map[c - 'A']++;
// 	PriorityQueue < Integer > queue = new PriorityQueue < > (26, Collections.reverseOrder());
// 	for (int f: map) {
// 		if (f > 0)
// 			queue.add(f);
// 	}
// 	int time = 0;
// 	while (!queue.isEmpty()) {
// 		int i = 0;
// 		List < Integer > temp = new ArrayList < > ();
// 		while (i <= n) {
// 			if (!queue.isEmpty()) {
// 				if (queue.peek() > 1)
// 					temp.add(queue.poll() - 1);
// 				else
// 					queue.poll();
// 			}
// 			time++;
// 			if (queue.isEmpty() && temp.size() == 0)
// 				break;
// 			i++;
// 		}
// 		for (int l: temp)
// 			queue.add(l);
// 	}
// 	return time;
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/task-scheduler/solution/ren-wu-diao-du-qi-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。