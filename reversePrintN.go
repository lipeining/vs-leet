package main

import (
	"fmt"
)

// func main() {
// 	//a := 0
// 	//fmt.Scan(&a)
// 	//fmt.Printf("%d\n", a)
// 	fmt.Printf("Hello World!\n")
// 	l := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val: 5,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	reversePrintReOrder(l, 2)
// 	reversePrintInOrder(l, 2)
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrintInOrder(head *ListNode, n int) {
	var dfs func(head *ListNode) int
	dfs = func(head *ListNode) int {
		if head == nil {
			return 0
		}
		ret := dfs(head.Next)
		fmt.Println("ret", ret, head.Val)
		if n == ret+1 {
			fmt.Println("form now:")
			fmt.Println(head.Val)
		}
		return ret + 1
	}
	dfs(head)
}
func reversePrintReOrder(head *ListNode, n int) {
	var dfs func(head *ListNode)
	dfs = func(head *ListNode) {
		if head == nil {
			return
		}
		dfs(head.Next)
		if n > 0 {
			fmt.Println(head.Val)
			n--
		}
	}
	dfs(head)
}
