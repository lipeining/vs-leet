/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
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
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{Val: -1}
	newHead.Next = head
	start := newHead
	for start != nil {
		keep := start.Next
		move := keep
		// 重复可能是重复 2，3，4，5 次
		for move != nil && move.Next != nil {
			if move.Val == move.Next.Val {
				move = move.Next
			} else {
				break
			}
		}
		if keep != move {
			// 说明 keep <-> move 是重复值
			// 考虑删除头结点
			start.Next = move.Next
		} else {
			start = keep
		}
	}
	return newHead.Next
}

// @lc code=end
