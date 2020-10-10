/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	mid := findMid(head)
	r := &TreeNode{Val: mid.Val}
	if mid == head {
		return r
	}
	r.Right = sortedListToBST(mid.Next)
	r.Left = sortedListToBST(head)
	return r
}
func findMid(head *ListNode) *ListNode {
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if prev != nil {
		prev.Next = nil
	}
	return slow
}
// func sortedListToBST(head *ListNode) *TreeNode {
// 	list := make([]int, 0)
// 	for head != nil {
// 		list = append(list, head.Val)
// 		head = head.Next
// 	}
// 	r := listToBST(list)
// 	return r
// }
// func listToBST(list []int) *TreeNode {
// 	length := len(list)
// 	if length == 0 {
// 		return nil
// 	}
// 	mid := length/2
// 	r := &TreeNode{Val:list[mid]}
// 	r.Left = listToBST(list[:mid])
// 	r.Right = listToBST(list[mid+1:])
// 	return r
// }
// @lc code=end

