/*
 * @lc app=leetcode.cn id=630 lang=golang
 *
 * [630] 课程表 III
 */

// @lc code=start
type pair struct{ t int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.t > b.t }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v pair)        { heap.Push(h, v) }
func (h *hp) pop() pair          { return heap.Pop(h).(pair) }
func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	queue := make(hp, 0)
	heap.Init(&queue)
	time := 0
	for _, c := range courses {
		if time+c[0] <= c[1] {
			queue.push(pair{c[0]})
			time += c[0]
		} else {
			if queue.Len() != 0 {
				top := queue.pop()
				if top.t > c[0] {
					time += c[0] - top.t
					queue.push(pair{c[0]})
				} else {
					heap.Push(&queue, top)
				}
			}
		}
	}
	return queue.Len()
}

// @lc code=end

