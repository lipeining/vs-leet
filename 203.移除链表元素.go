/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	pHead := &ListNode{
		-1,
		head,
	}
	p := pHead
	for p!=nil {
		n := p.Next
		if n != nil && n.Val == val {
			p.Next = n.Next
		} else {
			p = p.Next
		}
	}
	return pHead.Next
}
// @lc code=end

