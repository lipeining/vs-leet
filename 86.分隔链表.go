/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	// 遍历，分两个链表，官方思路
	beforeList := &ListNode{Val: -1}
	afterList := &ListNode{Val: -1}
	before := beforeList
	after := afterList
	for head != nil {
		// tmp := &ListNode{Val: head.Val}
		if head.Val >= x {
			// after.Next = tmp
			// after = after.Next
			after.Next = head
			after = after.Next
		} else {
			// before.Next = tmp
			// before = before.Next
			before.Next = head
			before = before.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterList.Next
	return beforeList.Next
}

// @lc code=end
