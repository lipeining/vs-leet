/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for  {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
	// slow, fast := head, head.Next
	// for fast != nil && fast.Next != nil {
	// 	if  slow == fast {
	// 		return true
	// 	}
	// 	slow  = slow.Next
	// 	if fast.Next != nil {
	// 		fast = fast.Next.Next
	// 	} else {
	// 		return false
	// 	}
	// }
	// return false
}

// @lc code=end

