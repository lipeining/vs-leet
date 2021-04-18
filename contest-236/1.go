package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	// findTheWinner(5, 2)
	// findTheWinner(6, 5)
	// findTheWinner(500, 2)
	// obj := Constructor(3, 1)
	// obj.AddElement(3)        // 当前元素为 [3]
	// obj.AddElement(1)        // 当前元素为 [3,1]
	// obj.CalculateMKAverage() // 返回 -1 ，因为 m = 3 ，但数据流中只有 2 个元素
	// obj.AddElement(10)       // 当前元素为 [3,1,10]
	// obj.CalculateMKAverage() // 最后 3 个元素为 [3,1,10]
	// // 删除最小以及最大的 1 个元素后，容器为 [3]
	// // [3] 的平均值等于 3/1 = 3 ，故返回 3
	// obj.AddElement(5)        // 当前元素为 [3,1,10,5]
	// obj.AddElement(5)        // 当前元素为 [3,1,10,5,5]
	// obj.AddElement(5)        // 当前元素为 [3,1,10,5,5,5]
	// obj.CalculateMKAverage() // 最后 3 个元素为 [5,5,5]
	// // 删除最小以及最大的 1 个元素后，容器为 [5]
	// // [5] 的平均值等于 5/1 = 5 ，故返回 5

	// ["MKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement","addElement","calculateMKAverage","addElement"]
	// [[30,10],[87859],[24884],[],[39764],[37246],[],[77016],[65453],[],[66229],[51558],[],[77202],[4526],[],[62945],[31817],[],[97483],[52991],[],[54305],[87130],[],[22677],[48120],[],[71933],[92149],[],[88407],[96760],[],[49114],[11334],[],[57536],[87001],[],[66641],[14147],[],[21457],[68281],[],[51545],[48566],[],[64186],[96046],[],[3877],[61515],[],[5700],[40440],[],[92194],[80585],[],[77750],[75783],[],[51590],[84825],[],[22329],[22098],[],[65830],[29746],[],[1613],[26152],[],[70729],[71872],[],[30432],[53013],[],[67342],[45066],[],[75733],[46305],[],[60180],[35295],[],[86405],[71827],[],[79816],[95604],[],[749],[50291],[],[97060]]
	// obj := Constructor(5, 2)
	// for i := 0; i <= 10; i++ {
	// 	obj.AddElement(i)
	// 	if i%5 == 0 {
	// 		obj.CalculateMKAverage()
	// 	}
	// }
	minSideJumps([]int{0, 1, 2, 3, 0})
	minSideJumps([]int{0, 1, 1, 3, 3, 0})
	minSideJumps([]int{0, 2, 1, 0, 3, 0})
}

// class Solution {
// 	public:
// 			int minSideJumps(vector<int>& ob) {
// 					int n = ob.size();
// 					vector<vector<int>> dp(n, vector<int>(3, 0x3f3f3f3f) );
// 					dp[0][1] = 0;
// 					dp[0][0] = 1;
// 					dp[0][2] = 1;

// 					for(int i=1; i<n; ++i)
// 					{
// 							if(ob[i] != 1) dp[i][0] = dp[i-1][0];
// 							if(ob[i] != 2) dp[i][1] = dp[i-1][1];
// 							if(ob[i] != 3) dp[i][2] = dp[i-1][2];
// 							if(ob[i] != 1) dp[i][0] = min(dp[i][0], min(dp[i][1], dp[i][2]) + 1);
// 							if(ob[i] != 2) dp[i][1] = min(dp[i][1], min(dp[i][0], dp[i][2]) + 1);
// 							if(ob[i] != 3) dp[i][2] = min(dp[i][2], min(dp[i][0], dp[i][1]) + 1);

// 					}
// 					return min(dp[n-1][0], min(dp[n-1][1], dp[n-1][2]));
// 			}
// 	};

// 	作者：zhouzzz
// 	链接：https://leetcode-cn.com/problems/minimum-sideway-jumps/solution/dp-by-zhouzzz-8tj2/
// 	来源：力扣（LeetCode）
// 	著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 4)
		for j := 0; j <= 3; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp[0][1] = 1
	dp[0][2] = 0
	dp[0][3] = 1
	for i := 1; i < n; i++ {
		if obstacles[i] != 1 {
			dp[i][1] = dp[i-1][1]
		}
		if obstacles[i] != 2 {
			dp[i][2] = dp[i-1][2]
		}
		if obstacles[i] != 3 {
			dp[i][3] = dp[i-1][3]
		}
		if obstacles[i] != 1 {
			dp[i][1] = min(dp[i][1], min(dp[i][2], dp[i][3])+1)
		}
		if obstacles[i] != 2 {
			dp[i][2] = min(dp[i][2], min(dp[i][1], dp[i][3])+1)
		}
		if obstacles[i] != 3 {
			dp[i][3] = min(dp[i][3], min(dp[i][1], dp[i][2])+1)
		}
	}
	fmt.Println(dp)
	ans := math.MaxInt32
	for i := 1; i <= 3; i++ {
		ans = min(ans, dp[n-1][i])
	}
	fmt.Println("ansssss", ans)
	return ans
}
func minSideJumpsArr(obstacles []int) int {
	n := len(obstacles) - 1
	dp := make([]int, 4)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp[0] = math.MaxInt32
	dp[2] = 0
	dp[1] = 1
	dp[3] = 1
	for i := 1; i < n; i++ {
		ndp := make([]int, 4)
		for j := 0; j < 4; j++ {
			ndp[j] = math.MaxInt32
		}
		for k := 1; k <= 3; k++ {
			if dp[k] == -1 || k == obstacles[i] {
				continue
			}
			ndp[k] = min(ndp[k], dp[k])
			for x := 1; x <= 3; x++ {
				if obstacles[i] != x {
					ndp[x] = min(dp[k]+1, ndp[x])
				}
			}
		}
		dp = ndp
	}
	fmt.Println(dp)
	ans := math.MaxInt32
	for i := 1; i <= 3; i++ {
		ans = min(ans, dp[i])
	}
	fmt.Println("ansssss", ans)
	return ans
}
func minSideJumps2(obstacles []int) int {
	n := len(obstacles)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 4)
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 到了第 i 个点，在  1, 2, 3 赛道上需要的最少 jumps
	dp[0][1] = 0
	dp[0][2] = 0
	dp[0][3] = 0
	for i := 1; i < n; i++ {
		if obstacles[i] == 1 {
			dp[i][1] = math.MaxInt32
			if obstacles[i-1] == 2 {
				dp[i][2] = min(dp[i-1][1]+1, dp[i-1][3]) + 1
				dp[i][3] = dp[i-1][3]
			} else if obstacles[i-1] == 3 {
				dp[i][3] = min(dp[i-1][1]+1, dp[i-1][2]) + 1
				dp[i][2] = dp[i-1][2]
			} else {
				dp[i][2] = dp[i-1][2]
				dp[i][3] = dp[i-1][3]
			}
		} else if obstacles[i] == 2 {
			dp[i][2] = math.MaxInt32
			if obstacles[i-1] == 1 {
				dp[i][1] = min(dp[i-1][2]+1, dp[i-1][3]) + 1
				dp[i][3] = dp[i-1][3]
			} else if obstacles[i-1] == 3 {
				dp[i][3] = min(dp[i-1][2]+1, dp[i-1][1]) + 1
				dp[i][1] = dp[i-1][1]
			} else {
				dp[i][1] = dp[i-1][1]
				dp[i][3] = dp[i-1][3]
			}
		} else if obstacles[i] == 3 {
			dp[i][3] = math.MaxInt32
			if obstacles[i-1] == 1 {
				dp[i][1] = min(dp[i-1][3]+1, dp[i-1][2]) + 1
				dp[i][2] = dp[i-1][2]
			} else if obstacles[i-1] == 2 {
				dp[i][2] = min(dp[i-1][3]+1, dp[i-1][1]) + 1
				dp[i][1] = dp[i-1][1]
			} else {
				dp[i][1] = dp[i-1][1]
				dp[i][2] = dp[i-1][2]
			}
		} else {
			// 没有障碍
			dp[i][1] = dp[i-1][1]
			dp[i][2] = dp[i-1][2]
			dp[i][2] = dp[i-1][3]
			if obstacles[i-1] == 1 {
				dp[i][1] = min(dp[i-1][2], dp[i-1][3]) + 1
			} else if obstacles[i-1] == 2 {
				dp[i][2] = min(dp[i-1][3], dp[i-1][1]) + 1
			} else if obstacles[i-1] == 3 {
				dp[i][3] = min(dp[i-1][2], dp[i-1][1]) + 1
			} else {
				// default
			}
		}
	}
	fmt.Println(dp)
	ans := math.MaxInt32
	for i := 1; i <= 3; i++ {
		ans = min(ans, dp[n-1][i])
	}
	fmt.Println("ansssss", ans)
	return ans
}

func findTheWinner(n int, k int) int {
	cnt := 0
	set := make([]bool, n+1)
	now := 1
	for cnt != n-1 {
		loop := 0
		for loop < k {
			if now <= n {
				if set[now] {
					// 已经出局
				} else {
					loop++
				}
				now++
			} else {
				now = 1
			}
		}
		// fmt.Println("now ", now, "set", now-1, " to true")
		cnt++
		if now == 1 {
			set[n] = true
		} else {
			set[now-1] = true
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if !set[i] {
			ans = i
			break
		}
	}
	fmt.Println("ansssss", ans)
	return ans
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

type MKAverage struct {
	maxHeap MaxHeap
	minHeap MinHeap
	list    []int
	index   int
	m       int
	k       int
}

func Constructor(m int, k int) MKAverage {
	maxHeap := make(MaxHeap, 0)
	minHeap := make(MinHeap, 0)
	heap.Init(&maxHeap)
	heap.Init(&minHeap)
	return MKAverage{
		maxHeap: maxHeap,
		minHeap: minHeap,
		list:    make([]int, 0),
		index:   0,
		m:       m,
		k:       k,
	}
}

func (this *MKAverage) AddElement(num int) {
	heap.Push(&this.maxHeap, &Item{
		value:    this.index,
		priority: num,
	})
	heap.Push(&this.minHeap, &Item{
		value:    this.index,
		priority: num,
	})
	this.list = append(this.list, num)
	this.index++
}

func (this *MKAverage) CalculateMKAverage() int {
	if this.index < this.m {
		return -1
	}
	sum := 0
	size := len(this.list)
	for i := 0; i < this.m; i++ {
		sum += this.list[size-i-1]
	}
	fmt.Println("base sumsums", this.list, sum)
	// 删除对应的 k * 2 个数
	big := 0
	small := 0
	bigAll := make(MaxHeap, 0)
	smallAll := make(MinHeap, 0)
	for big < this.k {
		top := heap.Pop(&this.maxHeap).(*Item)
		if top.value < size-this.m {
			// 已经是旧的 top  k 了
		} else {
			// fmt.Println("top big", top)
			sum -= top.priority
			big++
			// heap.Push(&this.maxHeap, top)
			bigAll = append(bigAll, top)
		}
	}
	for small < this.k {
		top := heap.Pop(&this.minHeap).(*Item)
		if top.value < size-this.m {
			// 已经是旧的 top  k 了
		} else {
			// fmt.Println("top small", top)
			sum -= top.priority
			small++
			// heap.Push(&this.minHeap, top)
			smallAll = append(smallAll, top)
		}
	}
	for _, top := range bigAll {
		heap.Push(&this.maxHeap, top)
	}
	for _, top := range smallAll {
		heap.Push(&this.minHeap, top)
	}
	ans := sum / (this.m - this.k*2)
	// 需要取整到最近的证书
	fmt.Println("anssssss", sum, ans)
	return ans
}

// func (this *MKAverage) AddElement(num int) {
// 	this.list = append(this.list, num)
// }

// func (this *MKAverage) CalculateMKAverage() int {
// 	size := len(this.list)
// 	if size < this.m {
// 		return -1
// 	}
// 	this.list = this.list[size-this.m : size]
// 	fmt.Println("after trim", this.list)
// 	arr := make([]int, this.m)
// 	copy(arr, this.list)
// 	sort.Ints(arr)
// 	sum := 0
// 	for i := 0; i < len(arr); i++ {
// 		sum += arr[i]
// 	}
// 	for i := 0; i < this.k; i++ {
// 		sum -= arr[i]
// 	}
// 	for j := 0; j < this.k; j++ {
// 		sum -= arr[len(arr)-1-j]
// 	}
// 	ans := sum / (len(arr) - this.k*2)
// 	fmt.Println("ansssss", ans)
// 	return ans
// }
