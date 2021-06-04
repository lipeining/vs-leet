/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
func reverseListNext(head *ListNode) *ListNode {
	// 递归不会
	// if head == nil || head.Next == nil {
	// 	return head
	// }
	// x := reverseList(head.Next)
	// x.Next = head
	// return x
// 	if (head == null || head.next == null) return head;
//     ListNode p = reverseList(head.next);
//     head.next.next = head;
//     head.next = null;
//     return p;

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/reverse-linked-list/solution/fan-zhuan-lian-biao-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	if head == nil || head.Next == nil {
		return head
	}
	// p := head.Next
	// head.Next = nil
	// for p != nil {
	// 	q := p.Next
	// 	p.Next = head
	// 	head = p
	// 	p = q
	// }
	// return head
	p := head
	var pre *ListNode = nil
	for p != nil {
		q := p.Next
		p.Next = pre
		pre = p
		p = q
	}
	return pre
}
// @lc code=end

