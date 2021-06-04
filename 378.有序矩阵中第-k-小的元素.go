/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第K小的元素
 */

// @lc code=start
type minitem struct{ t, i, j int }
type minheap []minitem

func (h minheap) Len() int            { return len(h) }
func (h minheap) Less(i, j int) bool  { a, b := h[i], h[j]; return a.t < b.t }
func (h minheap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minheap) Push(v interface{}) { *h = append(*h, v.(minitem)) }
func (h *minheap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *minheap) push(v minitem)     { heap.Push(h, v) }
func (h *minheap) pop() minitem       { return heap.Pop(h).(minitem) }
func (h *minheap) init()              { heap.Init(h) }

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	pq := make(minheap, 0)
	for i := 0; i < n; i++ {
		pq.push(minitem{t: matrix[i][0], i: i, j: 0})
	}
	for i := 0; i < k-1; i++ {
		top := pq.pop()
		if top.j != n-1 {
			pq.push(minitem{t: matrix[top.i][top.j+1], i: top.i, j: top.j + 1})
		}
	}
	top := pq[0]
	return top.t
}

// @lc code=end

