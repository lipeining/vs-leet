package main

import "fmt"

type TreeNode struct {
	Val  int
	list []*TreeNode
}

func answer(root *TreeNode) int {
	// bfs
	// if root == nil {
	// 	return 0
	// }
	// queue := make([]*TreeNode, 0)
	// queue = append(queue, root)
	// ans := 0
	// for len(queue) != 0 {
	// 	length := len(queue)
	// 	for i:=0;i<length;i++ {
	// 		node := queue[i]
	// 	}
	// 	queue = queue[:length]
	// }
	// return ans
	// dfs
	ans := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if isLeaf(root) {
			ans++
		} else {
			for _, node := range root.list {
				dfs(node)
			}
		}
	}
	dfs(root)
	return ans
}
func isLeaf(t *TreeNode) bool {
	length := len(t.list)
	return length == 0
}
func main() {
	root := &TreeNode{Val: 1}
	// 添加一层叶子节点 // 11, 12, 13
	// 给 11 添加 21 22 23
	for i := 1; i <= 3; i++ {
		tmp := &TreeNode{Val: i + 10}
		if i == 1 {
			for j := 1; j <= 3; j++ {
				leaf := &TreeNode{Val: j + 20}
				tmp.list = append(tmp.list, leaf)
			}
		}
		root.list = append(root.list, tmp)
	}
	ans := answer(root)
	fmt.Println(ans)
}
