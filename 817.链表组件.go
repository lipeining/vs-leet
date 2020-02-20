/*
 * @lc app=leetcode.cn id=817 lang=golang
 *
 * [817] 链表组件
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, G []int) int {
	gMap := make(map[int]bool)
	for _, n := range G {
		gMap[n] = true
	}
	ans := 0
	flag := false
	for head != nil {
		_, in := gMap[head.Val]
		if in {
			if flag {
				// do nothing
			} else {
				flag = true
			}
		} else {
			if flag {
				ans++
				flag = false
			} else {
				// do nothing
			}
		}
		// fmt.Println(flag, head.Val, ans)
		head = head.Next
	}
	if flag {
		ans++
	}
	return ans
}

// @lc code=end
