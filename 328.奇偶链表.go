/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	// 原地算法不会
	if head == nil || head.Next == nil {
		return head
	}
	odd := head
	even := head.Next
	o := odd
	e := even
	for e != nil {
		o.Next = e.Next
		if o.Next == nil {
			// 一共有偶数个节点
			break
		} else {
			o = o.Next
			e.Next = o.Next
			e = o.Next
		}
	}
	o.Next = even
	return odd
}

// @lc code=end
