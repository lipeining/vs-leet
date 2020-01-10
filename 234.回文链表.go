/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */
import "fmt"
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 借鉴别人的思路，快慢指针，找到中点，慢指针需要记录一个反转的链表
	slow, fast := head, head
	pre := head
	var prepare *ListNode= nil
	// pre.Next = nil
	// prepare := pre.Next
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
		pre.Next = prepare
		prepare = pre
	}
	if fast != nil {
		// 奇数时，
		slow = slow.Next
	}
	fmt.Println(slow, fast)
	// 重新对比对应的 pre(前半部分) slow（后半部分）
	for slow != nil && pre!=nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}
	return true

	// 反转的代码，需要套到 pre, prepare 上
	// p := head.Next
	// head.Next = nil
	// for p != nil {
	// 	q := p.Next
	// 	p.Next = head
	// 	head = p
	// 	p = q
	// }
	// return head
}
// @lc code=end

