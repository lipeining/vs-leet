package template

import "math"

// Prim算法循环|V|-1∣V∣−1次, 使用线性扫描算法寻找最小值的时间复杂度为O(|V|^2+|E|)O(∣V∣
// 	2
// 	 +∣E∣), 使用堆优化版Prim算法的时间复杂度是O(|E|log|V|)O(∣E∣log∣V∣).
func prim(n int, g [][]int) int {
	visited := make([]bool, n)
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[1] = 0
	ans := 0
	for i := 1; i <= n; i++ {
		node := -1
		for j := 1; j <= n; j++ {
			if !visited[j] && (node == -1 || dist[j] > dist[node]) {
				node = j
			}
		}
		if node == -1 {
			break
		}
		visited[node] = true
		ans += dist[node]
		for j := 1; j <= n; j++ {
			dist[j] = min(dist[j], g[node][j])
		}
	}
	return ans
}


// Hierholzer 算法过程：

// 选择任一顶点为起点，遍历所有相邻边。
// 深度搜索，访问相邻顶点。将经过的边都删除。
// 如果当前顶点没有相邻边，则将顶点入栈。
// 栈中的顶点倒序输出，就是从起点出发的欧拉回路。
