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
	s1, s2 := make([]int, 0), make([]int, 0)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	curry := 0
	var ans *ListNode
	for len(s1) != 0 || len(s2) != 0 || curry != 0 {
		var left int
		var right int
		if len(s1) != 0 {
			left = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) != 0 {
			right = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		sum := left + right + curry
		if sum >= 10 {
			curry = 1
			sum -= 10
		} else {
			curry = 0
		}
		node := &ListNode{Val: sum}
		node.Next = ans
		ans = node
	}
	return ans
}
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	// 采用自底向上的思路，某用户思路
// 	m := getLength(l1)
// 	n := getLength(l2)
// 	carry := 0
// 	var res *ListNode
// 	// 指针在函数之间是拷贝值，无法更新 head.Next and  res
// 	// 所以需要使用闭包
// 	var helper func(long *ListNode, short *ListNode, m, n int) int
// 	helper = func(long *ListNode, short *ListNode, m, n int) int {
// 		carry := 0
// 		longVal := 0
// 		shortVal := 0
// 		if m == n {
// 			if long == nil || short == nil {
// 				return 0
// 			}
// 			carry = helper(long.Next, short.Next, m, n)
// 			longVal = long.Val
// 			shortVal = short.Val
// 		} else {
// 			carry = helper(long.Next, short, m-1, n)
// 			longVal = long.Val
// 		}
// 		val := longVal + carry + shortVal
// 		head := &ListNode{Val: val % 10}
// 		head.Next = res
// 		res = head
// 		// fmt.Println(res, val)
// 		return val / 10
// 	}
// 	if m > n {
// 		carry = helper(l1, l2, m, n)
// 	} else if m < n {
// 		carry = helper(l2, l1, n, m)
// 	} else {
// 		carry = helper(l1, l2, m, n)
// 	}
// 	if carry != 0 {
// 		head := &ListNode{Val: 1}
// 		head.Next = res
// 		return head
// 	}
// 	return res
// }

func getLength(head *ListNode) int {
	n := 0
	for head != nil {
		n++
		head = head.Next
	}
	return n
}

// @lc code=end
