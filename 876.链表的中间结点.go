/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	len := getLen(head)
	mid := len / 2
	if len%2 == 1 {
		mid++
	}
	p := head
	fmt.Println(len, mid)
	for mid != 0 {
		mid--
		p = p.Next
	}
	return p
}
func getLen(head *ListNode) int {
	ans := 0
	p := head.Next
	for p != nil {
		ans++
		p = p.Next
	}
	return ans
}

// @lc code=end

