import "fmt"

/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// fmt.Println(head)
	// return head
	p := head.Next
	fakeHead := &ListNode{Val: head.Val}
	h := &ListNode{-1, fakeHead}
	// 还是有一个假的头结点比较简单
	for p != nil {
		t := h
		for t.Next != nil && t.Next.Val < p.Val {
			t = t.Next
		}
		oriNext := p.Next
		//
		tNext := t.Next
		t.Next = p
		p.Next = tNext

		p = oriNext
	}
	return h.Next
}

// @lc code=end
