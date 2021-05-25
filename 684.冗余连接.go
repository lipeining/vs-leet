/*
 * @lc app=leetcode.cn id=684 lang=golang
 *
 * [684] 冗余连接
 */

// @lc code=start
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	for i := 0; i <= n; i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x, y int) bool {
		x = find(x)
		y = find(y)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if !union(x, y) {
			return edge
		}
	}
	return []int{}
}

// @lc code=end

