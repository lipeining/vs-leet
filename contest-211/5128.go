package main

import "fmt"

func areConnected(n int, threshold int, queries [][]int) []bool {
	// create the graph
	// count the graph
	// check the queries
	G := NewGraph(n, 0)
	// 开始加边
	for _, edge := range MakeEdges(n, threshold) {
		G.addEdge(edge[0], edge[1])
	}
	return Connect(G, queries)
	// return nil
}

func MakeEdges(n, threshold int) [][]int {
	// 需要去重的边
	ans := make([][]int, 0)
	for z := 1; z <= n; z++ {
		if z > threshold {
			// 检查哪几个是可以整除 z 的
			for i := z + 1; i <= n; i++ {
				if i%z == 0 {
					ans = append(ans, []int{z, i})
				}
			}
		}
	}
	fmt.Println(ans)
	return ans
}

type Graph struct {
	V   int // 顶点数
	E   int // 边数
	Map map[int][]int
}

func NewGraph(v, e int) *Graph {
	return &Graph{V: v, E: e, Map: make(map[int][]int)}
}

func (G *Graph) addEdge(v, w int) {
	G.E++
	G.Map[v] = append(G.Map[v], w)
	G.Map[w] = append(G.Map[w], v)
}
func (G *Graph) adj(v int) []int {
	return G.Map[v]
}
func GetConnectCount(G *Graph) []int {
	// 从 1 开始
	marked := make([]bool, G.V+1)
	id := make([]int, G.V+1)
	count := 1
	var dfs func(node int)
	dfs = func(node int) {
		marked[node] = true
		id[node] = count
		for _, next := range G.adj(node) {
			if !marked[next] {
				dfs(next)
			}
		}
	}
	for i := 1; i <= G.V; i++ {
		if !marked[i] {
			dfs(i)
			count++
		}
	}
	return id
}
func Connect(G *Graph, queries [][]int) []bool {
	n := len(queries)
	ans := make([]bool, n)
	id := GetConnectCount(G)
	for i := 0; i < n; i++ {
		ans[i] = id[queries[i][0]] == id[queries[i][1]]
	}
	return ans
}
func test5128() {
	// 输入：n = 6, threshold = 2, queries = [[1,4],[2,5],[3,6]]
	// 输出：[false,false,true]
	// 解释：每个数的因数如下：
	// 1:   1
	// 2:   1, 2
	// 3:   1, 3
	// 4:   1, 2, 4
	// 5:   1, 5
	// 6:   1, 2, 3, 6
	// 所有大于阈值的的因数已经加粗标识，只有城市 3 和 6 共享公约数 3 ，因此结果是：
	// [1,4]   1 与 4 不连通
	// [2,5]   2 与 5 不连通
	// [3,6]   3 与 6 连通，存在路径 3--6

	fmt.Println(areConnected(6, 2, [][]int{{1, 4}, {2, 5}, {3, 6}}))

	// 	输入：n = 6, threshold = 0, queries = [[4,5],[3,4],[3,2],[2,6],[1,3]]
	// 输出：[true,true,true,true,true]
	// 解释：每个数的因数与上一个例子相同。但是，由于阈值为 0 ，所有的因数都大于阈值。因为所有的数字共享公因数 1 ，所以所有的城市都互相连通。

	fmt.Println(areConnected(6, 0, [][]int{{4, 5}, {3, 4}, {3, 2}, {2, 6}, {1, 3}}))

	// 输入：n = 5, threshold = 1, queries = [[4,5],[4,5],[3,2],[2,3],[3,4]]
	// 输出：[false,false,false,false,false]
	// 解释：只有城市 2 和 4 共享的公约数 2 严格大于阈值 1 ，所以只有这两座城市是连通的。
	// 注意，同一对节点 [x, y] 可以有多个查询，并且查询 [x，y] 等同于查询 [y，x] 。
	fmt.Println(areConnected(5, 1, [][]int{{4, 5}, {4, 5}, {3, 2}, {2, 3}, {3, 4}}))
}
