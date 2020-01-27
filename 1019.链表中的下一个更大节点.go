/*
 * @lc app=leetcode.cn id=1019 lang=golang
 *
 * [1019] 链表中的下一个更大节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	ans := make([]int, 0)
	if head == nil {
		return ans
	}
	p := head.Next
	for p!= nil {

	}
	return ans
}
// @lc code=end

