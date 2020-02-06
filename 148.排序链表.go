/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	// merge sort
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	nextHead := slow.Next
	slow.Next = nil
	prev := sortList(head)
	next := sortList(nextHead)
	return merge(prev, next)
}
func merge(prev *ListNode, next *ListNode) *ListNode {
	head := &ListNode{Val: -1}
	left := prev
	right := next
	p := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		p = p.Next
		// if left.Val < right.Val {
		// 	tmp := &ListNode{Val: left.Val}
		// 	p.Next = tmp
		// 	p = tmp
		// 	left = left.Next
		// } else {
		// 	tmp := &ListNode{Val: right.Val}
		// 	p.Next = tmp
		// 	p = tmp
		// 	right = right.Next
		// }
	}
	if left != nil {
		p.Next = left
	}
	if right != nil {
		p.Next = right
	}
	return head.Next
}

// @lc code=end
