/*
 * @lc app=leetcode.cn id=1046 lang=golang
 *
 * [1046] 最后一块石头的重量
 */

// @lc code=start
type maxitem struct{ t, i int }
type maxheap []maxitem

func (h maxheap) Len() int            { return len(h) }
func (h maxheap) Less(i, j int) bool  { a, b := h[i], h[j]; return a.t > b.t }
func (h maxheap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxheap) Push(v interface{}) { *h = append(*h, v.(maxitem)) }
func (h *maxheap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *maxheap) push(v maxitem)     { heap.Push(h, v) }
func (h *maxheap) pop() maxitem       { return heap.Pop(h).(maxitem) }
func (h *maxheap) init()              { heap.Init(h) }
func lastStoneWeight(stones []int) int {
	n := len(stones)
	pq := make(maxheap, n)
	for i := 0; i < n; i++ {
		pq[i] = maxitem{t: stones[i], i: i}
	}
	pq.init()
	for pq.Len() > 1 {
		y, x := pq.pop(), pq.pop()
		if y.t != x.t {
			pq.push(maxitem{t: y.t - x.t, i: -1})
		}
	}
	if pq.Len() != 0 {
		top := pq.pop()
		return top.t
	}
	return 0
}

// @lc code=end

