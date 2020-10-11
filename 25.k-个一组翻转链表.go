/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	p := pre
	for {
		cnt:=0
		for cnt<k && p.Next != nil {
			p = p.Next
			cnt++
		}
		if cnt != k {
			break
		}
		// p 为需要反转的最后一个节点
		newHead := p.Next
		p.Next = nil
		tmpHead := reverse(head)
		pre.Next = tmpHead
		head.Next = newHead
		pre = head
		p = pre
		head = newHead
	}
	return dummy.Next
}
func reverse(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
// @lc code=end

