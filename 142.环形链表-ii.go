/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// //  List   public ListNode detectCycle(ListNode head) {
// 	ListNode fast = head, slow = head;
// 	while (true) {
// 		if (fast == null || fast.next == null) return null;
// 		fast = fast.next.next;
// 		slow = slow.next;
// 		if (fast == slow) break;
// 	}
// 	fast = head;
// 	while (slow != fast) {
// 		slow = slow.next;
// 		fast = fast.next;
// 	}
// 	return fast;
// }

// 作者：jyd
// 链接：https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。授权，非商业转载请注明出处。
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

// @lc code=end
