/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
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

type minitem struct{ t, i int }
type minheap []minitem

func (h minheap) Len() int            { return len(h) }
func (h minheap) Less(i, j int) bool  { a, b := h[i], h[j]; return a.t < b.t }
func (h minheap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minheap) Push(v interface{}) { *h = append(*h, v.(minitem)) }
func (h *minheap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *minheap) push(v minitem)     { heap.Push(h, v) }
func (h *minheap) pop() minitem       { return heap.Pop(h).(minitem) }
func (h *minheap) init()              { heap.Init(h) }

type MedianFinder struct {
	lo maxheap
	hi minheap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		lo: make(maxheap, 0),
		hi: make(minheap, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.lo.push(maxitem{t: num, i: 0})
	this.hi.push(minitem{t: this.lo.pop().t, i: 0})
	if this.lo.Len() < this.hi.Len() {
		this.lo.push(maxitem{t: this.hi.pop().t, i: 0})
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		return float64(this.lo[0].t)
	}
	return (float64(this.lo[0].t) + float64(this.hi[0].t)) * 0.5
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

