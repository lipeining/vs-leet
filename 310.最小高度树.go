/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] 最小高度树
 */
package main

// import "fmt"

// func main() {
// 	findMinHeightTrees(1, [][]int{})
// 	findMinHeightTrees(2, [][]int{{1, 0}})
// 	findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}})
// 	findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}})
// 	// 59/68 time limit
// }

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	ans := make([]int, 0)
	if n <= 2 {
		for i := 0; i < n; i++ {
			ans = append(ans, i)
		}
		return ans
	}
	pre := make(map[int][]int)
	degree := make([]int, n)
	for _, edge := range edges {
		degree[edge[0]]++
		degree[edge[1]]++
		pre[edge[0]] = append(pre[edge[0]], edge[1])
		pre[edge[1]] = append(pre[edge[1]], edge[0])
	}
	// 将 出度为 1 的点进行删除
	queue := make([]int, 0)
	for index, de := range degree {
		if de == 1 {
			queue = append(queue, index)
		}
	}
	for n > 2 {
		size := len(queue)
		n -= size
		for i := 0; i < size; i++ {
			cur := queue[i]
			for _, next := range pre[cur] {
				degree[next]--
				if degree[next] == 1 {
					queue = append(queue, next)
				}
			}
		}
		queue = queue[size:]
	}
	for _, index := range queue {
		ans = append(ans, index)
	}
	// fmt.Println("ans", ans)
	return ans
	// curMin := n
	// for i := 0; i < n; i++ {
	// 	depth := getDepth(i, curMin, pre)
	// 	if depth < curMin {
	// 		curMin = depth
	// 		ans = make([]int, 0)
	// 		ans = append(ans, i)
	// 	} else if depth == curMin {
	// 		ans = append(ans, i)
	// 	}
	// }
	// fmt.Println("ans", curMin, ans)
	// return ans
}
func getDepth(root, curMin int, pre map[int][]int) int {
	level := 0
	queue := make([]int, 0)
	queue = append(queue, root)
	visited := make(map[int]bool)
	visited[root] = true
	for len(queue) != 0 {
		level++
		if level > curMin {
			return level
		}
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			for _, next := range pre[item] {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		queue = queue[size:]
	}
	// fmt.Println("level of root", root, "->", level)
	return level
}

// @lc code=end
