/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// if head == nil || head.Next == nil {
	// 	return head
	// }
	// left := head
	// right := left.Next
	// for right != nil {
	// 	left.Val, right.Val = right.Val, left.Val
	// 	left = right.Next
	// 	if left == nil || left.Next == nil {
	// 		break
	// 	}
	// 	right = left.Next
	// }
	// return head
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	p := head
	q := head.Next
	for p != nil && q != nil {
		pp := q.Next
		var qq *ListNode
		if pp != nil {
			qq = pp.Next
		}
		pre.Next = q
		q.Next = p
		p.Next = pp
		pre = p
		p = pp
		q = qq
	}
	return dummy.Next
}

// @lc code=end
