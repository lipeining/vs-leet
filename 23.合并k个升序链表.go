/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type pair struct{ node *ListNode }
 type hp []pair
 
 func (h hp) Len() int            { return len(h) }
 func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.node.Val < b.node.Val }
 func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
 func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
 func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
 func (h *hp) push(v pair)        { heap.Push(h, v) }
 func (h *hp) pop() pair          { return heap.Pop(h).(pair) }

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(hp, 0)
	dummy := &ListNode{Val: -1}
	for _, l := range lists {
		if l != nil {
			pq.push(pair{l})
		}
	}
	ptr := dummy
	for pq.Len() != 0 {
		top := pq.pop()
		node := top.node
		ptr.Next = node
		ptr = ptr.Next
		if node.Next != nil {
			pq.push(pair{node.Next})
		}
	}
	return dummy.Next
}
// @lc code=end

