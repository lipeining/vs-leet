import "fmt"

/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 采用自底向上的思路，某用户思路
	m := getLength(l1)
	n := getLength(l2)
	carry := 0
	var res *ListNode
	// 指针在函数之间是拷贝值，无法更新 head.Next and  res
	// 所以需要使用闭包
	var helper func(long *ListNode, short *ListNode, m, n int) int
	helper = func(long *ListNode, short *ListNode, m, n int) int {
		carry := 0
		longVal := 0
		shortVal := 0
		if m == n {
			if long == nil || short == nil {
				return 0
			}
			carry = helper(long.Next, short.Next, m, n)
			longVal = long.Val
			shortVal = short.Val
		} else {
			carry = helper(long.Next, short, m-1, n)
			longVal = long.Val
		}
		val := longVal + carry + shortVal
		head := &ListNode{Val: val % 10}
		head.Next = res
		res = head
		// fmt.Println(res, val)
		return val / 10
	}
	if m > n {
		carry = helper(l1, l2, m, n)
	} else if m < n {
		carry = helper(l2, l1, n, m)
	} else {
		carry = helper(l1, l2, m, n)
	}
	if carry != 0 {
		head := &ListNode{Val: 1}
		head.Next = res
		return head
	}
	return res
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
