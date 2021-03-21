package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// func1()
	func2()
	// func3()
	// func4()
	fmt.Println("anssssssssssssssss end")
}
func func1() {}
func maximumTime(time string) string {
	a, b, c, d := time[0], time[1], time[3], time[4]
	if a == '?' {
		a = '2'
	}
	if b == '?' {
		if a == '2' {
			b = '3'
		} else {
			b = '9'
		}
	}
	if c == '?' {
		c = '5'
	}
	if d == '?' {
		d = '9'
	}
	// fmt.Println("ansss")
	return string(a) + string(b) + ":" + string(c) + string(d)
}
func func2() {
	// 	输入：a = "aba", b = "caa"
	// 输出：2
	// 解释：满足每个条件的最佳方案分别是：
	// 1) 将 b 变为 "ccc"，2 次操作，满足 a 中的每个字母都小于 b 中的每个字母；
	// 2) 将 a 变为 "bbb" 并将 b 变为 "aaa"，3 次操作，满足 b 中的每个字母都小于 a 中的每个字母；
	// 3) 将 a 变为 "aaa" 并将 b 变为 "aaa"，2 次操作，满足 a 和 b 由同一个字母组成。
	// 最佳的方案只需要 2 次操作（满足条件 1 或者条件 3）。
	// 示例 2：

	// 输入：a = "dabadd", b = "cda"
	// 输出：3
	// 解释：满足条件 1 的最佳方案是将 b 变为 "eee" 。
	minCharacters("aba", "caa")
	minCharacters("dabadd", "cda")
}
func minCharacters(a string, b string) int {
	num := 26
	al, bl := make([]int, num), make([]int, num)
	for i := 0; i < len(a); i++ {
		al[a[i]-'a']++
	}
	for i := 0; i < len(b); i++ {
		bl[b[i]-'a']++
	}
	alp, blp := make([]int, num+1), make([]int, num+1)
	for i := 0; i < num; i++ {
		alp[i+1] = alp[i] + al[i]
		blp[i+1] = blp[i] + bl[i]
	}
	lena := len(a)
	lenb := len(b)
	ans := lena + lenb
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	fmt.Println(a, al, alp)
	fmt.Println(b, bl, blp)
	for i := 0; i < num; i++ {
		same := lena + lenb
		abig := lena + lenb
		bbig := lena + lenb
		aleft := alp[i] - alp[0]
		bleft := blp[i] - blp[0]
		aright := alp[num] - alp[i+1]
		bright := blp[num] - blp[i+1]
		if i == 25 || i == 0 {
			same = aleft + aright + bleft + bright
		} else {
			same = aleft + aright + bleft + bright
			abig = aleft + bright + bl[i]
			bbig = bleft + aright + al[i]
		}
		// fmt.Println("as char i ", i, "same is ", same, "abig is", abig, "bbig is", bbig)
		ans = min(ans, same)
		ans = min(ans, abig)
		ans = min(ans, bbig)
	}
	fmt.Println("anssssss", ans)
	return ans
}
func func3() {
	// 	输入：matrix = [[5,2],[1,6]], k = 1
	// 输出：7
	// 解释：坐标 (0,1) 的值是 5 XOR 2 = 7 ，为最大的值。
	// 示例 2：

	// 输入：matrix = [[5,2],[1,6]], k = 2
	// 输出：5
	// 解释：坐标 (0,0) 的值是 5 = 5 ，为第 2 大的值。
	// 示例 3：

	// 输入：matrix = [[5,2],[1,6]], k = 3
	// 输出：4
	// 解释：坐标 (1,0) 的值是 5 XOR 1 = 4 ，为第 3 大的值。
	// 示例 4：

	// 输入：matrix = [[5,2],[1,6]], k = 4
	// 输出：0
	// 解释：坐标 (1,1) 的值是 5 XOR 2 XOR 1 XOR 6 = 0 ，为第 4 大的值。
	kthLargestValue([][]int{{5, 2}, {1, 6}}, 1)
	kthLargestValue([][]int{{5, 2}, {1, 6}}, 2)
	kthLargestValue([][]int{{5, 2}, {1, 6}}, 3)
	kthLargestValue([][]int{{5, 2}, {1, 6}}, 4)
}

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
func kthLargestValue(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
			} else if i == 0 {
				matrix[0][j] ^= matrix[0][j-1]
			} else if j == 0 {
				matrix[i][0] ^= matrix[i-1][0]
			} else {
				matrix[i][j] ^= matrix[i-1][j-1] ^ matrix[i-1][j] ^ matrix[i][j-1]
			}
			priority := matrix[i][j]
			item := &Item{
				value:    priority,
				priority: priority,
				index:    i*m + j,
			}
			heap.Push(&pq, item)
			// for k := 0; k < len(pq); k++ {
			// 	fmt.Println(pq[k])
			// }
			// fmt.Println("eneneenen pq")
			if pq.Len() > k {
				heap.Pop(&pq)
			}
		}
	}
	fmt.Println(matrix)
	ans := heap.Pop(&pq).(*Item).value
	fmt.Println("anssss", ans)
	return ans
}
func func4() {}
