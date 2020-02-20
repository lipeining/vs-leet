/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1}
	plus := 0
	p := head
	for l1 != nil || l2 != nil {
		leftVal := 0
		if l1 != nil {
			leftVal = l1.Val
			l1 = l1.Next
		}
		rightVal := 0
		if l2 != nil {
			rightVal = l2.Val
			l2 = l2.Next
		}
		curr := leftVal + rightVal + plus
		if curr >= 10 {
			curr -= 10
			plus = 1
		} else {
			plus = 0
		}
		tmp := &ListNode{Val: curr}
		p.Next = tmp
		p = p.Next
	}
	if plus != 0 {
		tmp := &ListNode{Val: plus}
		p.Next = tmp
	}
	return head.Next
}

// @lc code=end
