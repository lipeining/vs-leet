import "fmt"

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
type node struct {
	index int
	num   int
}

func nextLargerNodes(head *ListNode) []int {
	ans := make([]int, 10000)
	if head == nil {
		return ans
	}
	p := head
	index := 0
	stack := make([]*node, 0)
	for p != nil {
		length := len(stack)
		if length == 0 {
			stack = append(stack, &node{index: index, num: p.Val})
		} else {
			flag := false
			for len(stack) != 0 {
				top := stack[len(stack)-1]
				if p.Val > top.num {
					stack = stack[:len(stack)-1]
					ans[top.index] = p.Val
				} else {
					stack = append(stack, &node{index: index, num: p.Val})
					flag = true
					break
				}
			}
			if !flag {
				stack = append(stack, &node{index: index, num: p.Val})
			}
		}
		// fmt.Println(stack, ans[:index+1])
		p = p.Next
		index++
	}
	// fmt.Println(stack)
	return ans[:index]
}

// @lc code=end
