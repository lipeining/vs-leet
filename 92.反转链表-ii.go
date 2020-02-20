/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	curr := 1
	newHead := &ListNode{Val: -1}
	newHead.Next = head
	p := head
	for curr < m {
		p = p.Next
		curr++
	}
	start := p
	list := make([]int, 0)
	for curr <= n {
		list = append(list, p.Val)
		p = p.Next
		curr++
	}
	// end := p.Next
	for i := len(list) - 1; i >= 0; i-- {
		start.Val = list[i]
		start = start.Next
	}
	return newHead.Next
}

// @lc code=end
