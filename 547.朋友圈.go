/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 朋友圈
 */

// @lc code=start
// func findCircleNum(M [][]int) int {
// 	var dfs func(M [][]int, visited []bool, start int)
// 	dfs = func(M [][]int, visited []bool, start int) {
// 		for j:=0;j<len(M);j++ {
// 			if M[start][j] == 1 && !visited[j] {
// 				visited[j] = true
// 				dfs(M, visited, j)
// 			}
// 		}
// 	}
// 	visited := make([]bool, len(M))
// 	ans := 0
// 	for j:=0;j<len(M);j++ {
// 		if !visited[j] {
// 			dfs(M, visited, j)
// 			ans++
// 		}
// 	}
// 	return ans
// }

// ---------------------------
// func find(parent []int, i int) int {
// 	if parent[i] == -1 {
// 		return i
// 	}
// 	return find(parent, parent[i])
// }
// func union(parent []int, x int, y int) {
// 	xset := find(parent, x)
// 	yset := find(parent, y)
// 	if xset != yset {
// 		parent[xset] = yset
// 	}
// }
// func findCircleNum(M [][]int) int {
// 	parent := make([]int, len(M))
// 	for i:=0;i<len(parent);i++ {
// 		parent[i] = -1
// 	}
// 	for i:=0;i<len(M);i++ {
// 		for j:=0;j<len(M);j++ {
// 			if i!=j && M[i][j] == 1 {
// 				union(parent, i, j)
// 			}
// 		}
// 	}
// 	count := 0
// 	for i:=0;i<len(parent);i++ {
// 		if parent[i] == -1 {
// 			count++
// 		}
// 	}
// 	fmt.Println(parent)
// 	return count
// }
func find(parent map[int]int, i int) int {
	if parent[i] == -1 {
		return i
	}
	return find(parent, parent[i])
}
func union(parent map[int]int, x int, y int) {
	xset := find(parent, x)
	yset := find(parent, y)
	if xset != yset {
		parent[xset] = yset
	}
}
func findCircleNum(M [][]int) int {
	parent := make(map[int]int)
	for i:=0;i<len(M);i++ {
		parent[i] = -1
	}
	for i:=0;i<len(M);i++ {
		for j:=0;j<len(M);j++ {
			if i!=j && M[i][j] == 1 {
				union(parent, i, j)
			}
		}
	}
	count := 0
	for _,v := range parent {
		if v == -1 {
			count++
		}
	}
	fmt.Println(parent)
	return count
}
// @lc code=end
	// // 想法是：邻接矩阵的深度遍历吗？
	// N:=len(M)
	// if N==0 {
	// 	return 0
	// }
	// cycle := make(map[int]map[int]bool)
	// for i:=0;i<N;i++ {
	// 	for j:=0;j<N;j++ {
	// 		if i!=j && M[i][j] == 1 {
	// 			sub,ok := cycle[i]
	// 			if !ok {
	// 				sub = make(map[int]bool)
	// 			}
	// 			sub[j] = true
	// 			cycle[i] = sub
	// 		}
	// 	}
	// }
	// fmt.Println(cycle)
	// visited := make([]bool, N)
	// var dfs func(cycle map[int]map[int]bool, visited []bool, start int)
	// dfs = func(cycle map[int]map[int]bool, visited []bool, start int) {
	// 	if visited[start] {
	// 		return
	// 	}
	// 	visited[start] = true
	// 	for k,_ := range cycle[start] {
	// 		dfs(cycle, visited, k)
	// 	}
	// }
	// dfs(cycle, visited, 0)
	// // 如何通过 visited 判断是否需要新的一轮循环
	// return 0
