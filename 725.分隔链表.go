/*
 * @lc app=leetcode.cn id=725 lang=golang
 *
 * [725] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
	length := getLength(root)
	ans := make([]*ListNode, k)
	mod := length % k
	every := length / k
	if every == 0 {
		every = 1
		mod = 0
	}
	p := root
	used := 0
	for i := 0; i < k; i++ {
		list := p
		counter := 1
		modLeft := 0
		if mod > 0 {
			modLeft = 1
			mod--
		}
		if used >= length {
			continue
		}
		for counter < every+modLeft && used+counter < length {
			p = p.Next
			counter++
		}
		next := p.Next
		p.Next = nil
		p = next
		ans[i] = list
		used += counter
	}
	return ans
}
func getLength(head *ListNode) int {
	n := 0
	for head != nil {
		n++
		head = head.Next
	}
	return n
}

// @lc code=end
