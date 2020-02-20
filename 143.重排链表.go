/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	// 官方简单的一个思路
	list := make([]*ListNode, 0)
	p := head
	for p != nil {
		list = append(list, p)
		p = p.Next
	}
	length := len(list)
	i := 0
	j := length - 1
	for i < j {
		list[i].Next = list[j]
		i++
		if i == j {
			break
		}
		list[j].Next = list[i]
		j--
	}
	list[i].Next = nil
}

// @lc code=end
