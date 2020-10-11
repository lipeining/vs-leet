/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	p := dummy
	for !breakList(lists) {
		index := findMinIndex(lists)
		node := &ListNode{Val: lists[index].Val}
		p.Next = node
		p = p.Next
		lists[index] = lists[index].Next
	}
	return dummy.Next
}
func findMinIndex(lists []*ListNode) int {
	index := 0
	min := math.MaxInt32
	for i,node:=range lists {
		if node != nil && node.Val < min {
			min = node.Val
			index = i
		}
	}	
	return index
}
func breakList(lists []*ListNode) bool {
	cnt := 0
	for _,node := range lists {
		if node == nil {
			cnt++
		}
	}
	return cnt == len(lists)
}
// @lc code=end

