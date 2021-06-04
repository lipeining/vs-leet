/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
type pair struct{ t, i int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.t > b.t }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v pair)        { heap.Push(h, v) }
func (h *hp) pop() pair          { return heap.Pop(h).(pair) }
func (h *hp) init()              { heap.Init(h) }

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	pq := make(hp, 0)
	ans := make([]int, 0)
	for i := 0; i < n && i < k; i++ {
		pq.push(pair{t: nums[i], i: i})
	}
	top := pq[0]
	ans = append(ans, top.t)
	for i := k; i < n; i++ {
		pq.push(pair{t: nums[i], i: i})
		top = pq[0]
		for top.i <= i-k {
			pq.pop()
			top = pq[0]
		}
		ans = append(ans, top.t)
	}
	return ans
}

// @lc code=end

