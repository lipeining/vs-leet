/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{-1, nil}
	p := head
	for ;l1 != nil || l2 != nil; {
		if l1 == nil {
			tmp := &ListNode{l2.Val, nil}
			p.Next = tmp
			p = p.Next
			l2 = l2.Next
			continue
		}
		if l2 == nil {
			tmp := &ListNode{l1.Val, nil}
			p.Next = tmp
			p = p.Next
			l1 = l1.Next
			continue
		}
		v1 := l1.Val
		v2 := l2.Val
		if v1 < v2 {
			tmp := &ListNode{l1.Val, nil}
			p.Next = tmp
			p = p.Next
			l1 = l1.Next
		} else {
			tmp := &ListNode{l2.Val, nil}
			p.Next = tmp
			p = p.Next
			l2 = l2.Next
		}
	}
	return head.Next
}
// @lc code=end

