package main

import (
	"fmt"
)

func main() {
	{
		n := 7
		edges := [][]int{
			{0, 1},
			{0, 2},
			{1, 4},
			{1, 5},
			{2, 3},
			{2, 6},
		}
		labels := "abaedcd"
		fmt.Println(countSubTrees(n, edges, labels))
	}
	{
		// n = 4, edges = [[0,1],[1,2],[0,3]], labels = "bbbb"
		n := 4
		edges := [][]int{
			{0, 1},
			{1, 2},
			{0, 3},
		}
		labels := "bbbb"
		fmt.Println(countSubTrees(n, edges, labels))
	}

	{
		// 4
		// [[0,2],[0,3],[1,2]]
		// "aeed"
		// want [1,1,2,1]
		n := 4
		edges := [][]int{
			{0, 2},
			{0, 3},
			{1, 2},
		}
		labels := "aeed"
		fmt.Println(countSubTrees(n, edges, labels))
	}
}
func countSubTrees(n int, edges [][]int, labels string) []int {
	ans := make([]int, n)
	if n == 0 {
		return ans
	}
	for i := 0; i < n; i++ {
		visited := make([]bool, n)
		ans[i] = dfs(n, edges, labels, visited, i, labels[i])
	}
	return ans
}
func dfs(n int, edges [][]int, labels string, visited []bool, head int, want byte) int {
	num := 0
	for i, edge := range edges {
		if visited[i] {
			continue
		}
		visited[i] = true
		if edge[0] == head {
			if labels[edge[0]] == want {
				num++
			}
			num += dfs(n, edges, labels, visited, edge[1], want)
		} else if edge[1] == head {
			if labels[edge[1]] == want {
				num++
			}
			num += dfs(n, edges, labels, visited, edge[0], want)
		}
	}
	return num
}

// func countSubTrees(n int, edges [][]int, labels string) []int {
// 	ans := make([]int, n)
// 	if n == 0 {
// 		return ans
// 	}
// 	visited := make([]bool, len(edges))
// 	root := buildTree(n, edges, visited, labels, 0)
// 	// fmt.Println(root)
// 	for i := 0; i < n; i++ {
// 		head := findHead(root, i)
// 		count := countNo(head, labels[i])
// 		ans[i] = count
// 	}
// 	return ans
// }

// type tree struct {
// 	index    int
// 	val      byte
// 	children []*tree
// }

// func findHead(root *tree, index int) *tree {
// 	if root == nil {
// 		return nil
// 	}
// 	if root.index == index {
// 		return root
// 	}
// 	for _, sub := range root.children {
// 		if ret := findHead(sub, index); ret != nil {
// 			return ret
// 		}
// 	}
// 	return nil
// }
// func countNo(root *tree, want byte) int {
// 	if root == nil {
// 		return 0
// 	}
// 	current := 0
// 	if root.val == want {
// 		current = 1
// 	}
// 	for _, sub := range root.children {
// 		current += countNo(sub, want)
// 	}
// 	return current
// }
// func buildTree(n int, edges [][]int, visited []bool, labels string, head int) *tree {
// 	root := tree{val: labels[head], index: head}
// 	for i := 0; i < len(edges); i++ {
// 		if visited[i] {
// 			continue
// 		}
// 		if edges[i][0] == head {
// 			visited[i] = true
// 			root.children = append(root.children, buildTree(n, edges, visited, labels, edges[i][1]))
// 		} else if edges[i][1] == head {
// 			visited[i] = true
// 			root.children = append(root.children, buildTree(n, edges, visited, labels, edges[i][0]))
// 		}
// 	}
// 	return &root
// }
