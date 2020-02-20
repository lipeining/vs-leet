/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	// 需要考虑 K 大于链表长度的情况
	length := getLen(head)
	// fmt.Println(k, length, k%length)
	n := k % length
	if n == 0 {
		return head
	}
	slow, fast := head, head
	for n != 0 {
		fast = fast.Next
		n--
	}
	// if fast == nil {
	// 	// 删除头结点吧
	// 	return slow.Next
	// }
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	newHead := slow.Next
	fast.Next = head
	slow.Next = nil
	return newHead
}
func getLen(head *ListNode) int {
	l := 0
	for head != nil {
		l++
		head = head.Next
	}
	return l
}

// @lc code=end
