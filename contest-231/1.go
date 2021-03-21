package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// func1()
	// func2()
	func3()
}
func func1() {
	fmt.Println("anss", checkOnesSegment("1001"))
	fmt.Println("anss", checkOnesSegment("110"))
	fmt.Println("anss", checkOnesSegment("1"))
	fmt.Println("anss", checkOnesSegment("10"))
}
func checkOnesSegment(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	i := 0
	for i < n {
		if s[i] == '0' {
			break
		}
		i++
	}
	// 之后不会有 1
	for i < n {
		if s[i] == '1' {
			return false
		}
		i++
	}
	return true
}
func func2() {
	// 输入：nums = [1,-1,1], limit = 3, goal = -4
	// 输出：2
	// 解释：可以将 -2 和 -3 添加到数组中，数组的元素总和变为 1 - 1 + 1 - 2 - 3 = -4 。
	// 示例 2：

	// 输入：nums = [1,-10,9,1], limit = 100, goal = 0
	// 输出：1
	minElements([]int{1, -1, 1}, 3, -4)
	minElements([]int{1, -10, 9, 1}, 100, 0)
	// 	[-1,0,1,1,1]
	// 1
	// 771843707
	minElements([]int{-1, 0, 1, 1, 1}, 1, 771843707)
}
func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum == goal {
		return 0
	}
	// 贪心算法
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	need := abs(goal - sum)
	cnt := need / limit
	if need%limit != 0 {
		cnt++
	}
	// for now != need {
	// 	a := abs(now - need)
	// 	cur := min(a, limit)
	// 	cnt++
	// 	now += cur
	// 	fmt.Println("cnt", cnt, "cur", cur)
	// }
	fmt.Println("ansss", cnt)
	return cnt
}
func func3() {
	// 输入：n = 5, edges = [[1,2,3],[1,3,3],[2,3,1],[1,4,2],[5,2,2],[3,5,1],[5,4,10]]
	// 输出：3
	countRestrictedPaths(5, [][]int{
		{1, 2, 3},
		{1, 3, 3},
		{2, 3, 1},
		{1, 4, 2},
		{5, 2, 2},
		{3, 5, 1},
		{5, 4, 10},
	})
	// 输入：n = 7, edges = [[1,3,1],[4,1,2],[7,3,4],[2,5,3],[5,6,1],[6,7,2],[7,5,3],[2,6,4]]
	// 输出：1
	// 解释：每个圆包含黑色的节点编号和蓝色的 distanceToLastNode 值。唯一一条受限路径是：1 --> 3 --> 7 。
	countRestrictedPaths(7, [][]int{
		{1, 3, 1},
		{4, 1, 2},
		{7, 3, 4},
		{2, 5, 3},
		{5, 6, 1},
		{6, 7, 2},
		{7, 5, 3},
		{2, 6, 4},
	})
	// 6
	// [[6,2,9],[2,1,7],[6,5,10],[4,3,1],[3,1,8],[4,6,4],[5,1,7],[1,4,7]
	// ,[4,5,3],[3,6,6],[5,3,9],[1,6,6],[3,2,2],[5,2,8]]
	//4
	countRestrictedPaths(6, [][]int{
		{6, 2, 9},
		{2, 1, 7},
		{6, 5, 10},
		{4, 3, 1},
		{3, 1, 8},
		{4, 6, 4},
		{5, 1, 7},
		{1, 4, 7},
		{4, 5, 3},
		{3, 6, 6},
		{5, 3, 9},
		{1, 6, 6},
		{3, 2, 2},
		{5, 2, 8},
	})
}
func countRestrictedPaths(n int, edges [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	t := make(map[int]map[int]int)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if t[u] == nil {
			t[u] = make(map[int]int)
		}
		t[u][v] = w
		if t[v] == nil {
			t[v] = make(map[int]int)
		}
		t[v][u] = w
	}
	// fmt.Println(t)
	dis := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dis[i] = math.MaxInt32
	}
	var dfs func(end int, now int, path int, visited []bool)
	dfs = func(end int, now int, path int, visited []bool) {
		if now == end {
			dis[end] = min(dis[end], path)
			return
		}
		for v, weight := range t[now] {
			if !visited[v] {
				visited[v] = true
				dfs(end, v, path+weight, visited)
				visited[v] = false
			}
		}
	}
	for end := n; end >= 1; end-- {
		visited := make([]bool, n+1)
		visited[n] = true
		dfs(end, n, 0, visited)
	}
	// fmt.Println("dissss", dis)
	mod := int(1e9 + 7)
	ways := make([]int, n+1)
	ways[1] = 1
	var dfsway func(now int) int
	dfsway = func(now int) int {
		if ways[now] != 0 {
			return ways[now]
		}
		cur := 0
		for p := range t[now] {
			if dis[p] > dis[now] {
				pway := dfsway(p)
				fmt.Println("now", now, "from p", p, "pway", pway, "cur", cur)
				cur += pway
			}
		}
		ways[now] = cur
		return cur
	}
	dfsway(n)
	ans := ways[n]
	ans %= mod
	fmt.Println("anssssssssssssssssss", ways, ans)
	return ans
}
func countRestrictedPathsMap(n int, edges [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	t := make(map[int]map[int]int)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if t[u] == nil {
			t[u] = make(map[int]int)
		}
		t[u][v] = w
		if t[v] == nil {
			t[v] = make(map[int]int)
		}
		t[v][u] = w
	}
	// fmt.Println(t)
	dis := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dis[i] = math.MaxInt32
	}
	var dfs func(end int, now int, path int, visited []bool)
	dfs = func(end int, now int, path int, visited []bool) {
		if now == end {
			dis[end] = min(dis[end], path)
			return
		}
		for v, weight := range t[now] {
			if !visited[v] {
				visited[v] = true
				dfs(end, v, path+weight, visited)
				visited[v] = false
			}
		}
	}
	for end := n; end >= 1; end-- {
		visited := make([]bool, n+1)
		visited[n] = true
		dfs(end, n, 0, visited)
	}
	// fmt.Println("dissss", dis)
	mod := int(1e9 + 7)
	ways := make([]int, n+1)
	ways[1] = 1
	queue := make(map[int]bool)
	queue[1] = true
	for len(queue) != 0 {
		// fmt.Println("queue", queue)
		rm := make([]int, 0)
		next := make([]int, 0)
		for now := range queue {
			for v := range t[now] {
				if dis[now] > dis[v] {
					ways[v] += ways[now]
					// fmt.Println("on now", now, "to v", v, ways[now], ways[v])
					if !queue[v] {
						// queue[v] = true
						next = append(next, v)
					}
				}
			}
			rm = append(rm, now)
		}
		for _, node := range rm {
			delete(queue, node)
		}
		for _, node := range next {
			// delete(queue, node)
			queue[node] = true
		}
	}
	ans := ways[n]
	ans %= mod
	fmt.Println("anssssssssssssssssss", ways, ans)
	return ans
}
func countRestrictedPathsArr(n int, edges [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	t := make(map[int]map[int]int)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if t[u] == nil {
			t[u] = make(map[int]int)
		}
		t[u][v] = w
		if t[v] == nil {
			t[v] = make(map[int]int)
		}
		t[v][u] = w
	}
	fmt.Println(t)
	dis := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dis[i] = math.MaxInt32
	}
	var dfs func(end int, now int, path int, visited []bool)
	dfs = func(end int, now int, path int, visited []bool) {
		if now == end {
			dis[end] = min(dis[end], path)
			return
		}
		for v, weight := range t[now] {
			if !visited[v] {
				visited[v] = true
				dfs(end, v, path+weight, visited)
				visited[v] = false
			}
		}
	}
	for end := n; end >= 1; end-- {
		visited := make([]bool, n+1)
		visited[n] = true
		dfs(end, n, 0, visited)
	}
	fmt.Println("dissss", dis)
	mod := int(1e9 + 7)
	ways := make([]int, n+1)
	ways[1] = 1
	queue := make([]int, 0)
	queue = append(queue, 1)
	for len(queue) != 0 {
		size := len(queue)
		next := make(map[int]bool)
		for i := 0; i < size; i++ {
			now := queue[i]
			for v := range t[now] {
				if dis[now] > dis[v] {
					ways[v] += ways[now]
					fmt.Println("on now", now, "to v", v, ways[now], ways[v])
					// queue = append(queue, v)
					next[v] = true
				}
			}
		}
		for i := 0; i < size; i++ {
			delete(next, queue[i])
		}
		queue = queue[size:]
		for node := range next {
			queue = append(queue, node)
		}
		fmt.Println("before sort", queue)
		sort.Slice(queue, func(l, r int) bool {
			return dis[l] > dis[r]
		})
		fmt.Println("after sort", queue)
	}
	ans := ways[n]
	ans %= mod
	fmt.Println("ansssssssssssssssssssssss", ways, ans)
	return ans
}
func func4() {}
