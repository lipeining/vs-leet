/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// 递归 或者 判断一下
	p := head
	for p != nil {
		n := p.Next
		if n != nil {
			if p.Val == n.Val {
				p.Next = n.Next
				continue
			}
		}
		p = p.Next
	}
	return head
}
// @lc code=end

