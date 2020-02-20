/*
 * @lc app=leetcode.cn id=1171 lang=golang
 *
 * [1171] 从链表中删去总和值为零的连续节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	// 两次相同的前缀和，别人的思路
	fakeHead := &ListNode{0, head}
	m := make(map[int]*ListNode)
	p := fakeHead
	sum := 0
	for p != nil {
		sum += p.Val
		m[sum] = p
		p = p.Next
	}
	sum = 0
	p = fakeHead
	for p != nil {
		sum += p.Val
		if m[sum] != p {
			p.Next = m[sum].Next
		}
		p = p.Next
	}
	return fakeHead.Next
}

// @lc code=end
