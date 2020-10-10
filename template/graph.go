package template

type Graph struct {
	V   int // 顶点数
	E   int // 边数
	Map map[int][]int
	Dis [][]int
}

func NewGraph(v, e int) *Graph {
	dis := make([][]int, v)
	for i := 0; i < v; i++ {
		dis[i] = make([]int, v)
	}
	return &Graph{V: v, E: e, Map: make(map[int][]int), Dis: dis}
}

func (G *Graph) addEdge(v, w int) {
	G.E++
	G.Map[v] = append(G.Map[v], w)
	G.Map[w] = append(G.Map[w], v)
}
func (G *Graph) adj(v int) []int {
	return G.Map[v]
}

func DepthFirstSearch(G *Graph, start int) {
	marked := make([]bool, G.V)
	count := 0 // 连接 start 的点数
	var dfs func(v int)
	dfs = func(v int) {
		if marked[v] {
			return
		}
		// fmt.Println(v)
		count++
		marked[v] = true
		for _, w := range G.adj(v) {
			if !marked[w] {
				dfs(w)
			}
		}
	}
	dfs(start)
	// return count
}

func PathTo(G *Graph, start, end int) []int {
	marked := make([]bool, G.V)
	edgesTo := make([]int, G.V) // 从 A 到 B 的记录
	var dfs func(v int)
	dfs = func(v int) {
		if marked[v] {
			return
		}
		marked[v] = true
		for _, w := range G.adj(v) {
			if !marked[w] {
				edgesTo[w] = v
				dfs(w)
			}
		}
	}
	dfs(start)
	if !marked[end] {
		return nil
	}
	s := end
	path := make([]int, 0)
	for s != start {
		path = append(path, s)
		s = edgesTo[s]
	}
	path = append(path, start)
	return path
}

func TOPO(G *Graph) []int {
	marked := make([]bool, G.V)
	reverse := make([]int, 0)
	var dfs func(v int)
	dfs = func(v int) {
		if marked[v] {
			return
		}
		marked[v] = true
		for _, w := range G.adj(v) {
			if !marked[w] {
				dfs(w)
			}
		}
		reverse = append(reverse, v)
	}
	for i := 0; i < G.V; i++ {
		if !marked[i] {
			dfs(i)
		}
	}
	return reverse
}

func BreadthFirstSearch(G *Graph, start int, end int) []int {
	marked := make([]bool, G.V)
	edgesTo := make([]int, G.V) // 从 A 到 B 的记录
	queue := make([]int, 0)
	queue = append(queue, start)
	marked[start] = true
	for len(queue) != 0 {
		top := queue[0]
		for _, w := range G.adj(top) {
			if !marked[w] {
				edgesTo[w] = top
				marked[w] = true
				queue = append(queue, w)
			}
		}
		queue = queue[1:]
	}
	s := end
	path := make([]int, 0)
	for s != start {
		path = append(path, s)
		s = edgesTo[s]
	}
	path = append(path, start)
	return path
}

func Connect(G *Graph, v, w int) bool {
	marked := make([]bool, G.V)
	id := make([]int, G.V)
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
	for i := 0; i < G.V; i++ {
		if !marked[i] {
			dfs(i)
			count++
		}
	}
	return id[v] == id[w]
}

// 一个连通图，若干节点，节点可能有数值，但是路径一定有权值。并且路径不能为负。否则Dijkstra就不适用。
// Dijkstra算法的前提：
// 首先，Dijkstra处理的是带正权值的有权图，那么，就需要一个二维数组（如果空间大用list数组）存储各个点到达(边)的权值大小。(邻接矩阵或者邻接表存储)
// 其次，还需要一个boolean数组判断那些点已经确定最短长度，那些点没有确定。int数组记录距离(在算法执行过程可能被多次更新)。
// 需要优先队列加入已经确定点的周围点。每次抛出确定最短路径的那个并且确定最短，直到所有点路径确定最短为止。
// 我们用一个数组dis记录源点s到其他点的最短距离，起始时dis[s] = 0，其他值设为无穷大
// 我们从未访问过的点当中选择距离最小的点u，将它标记为已访问
// 遍历u所有可以连通的点v，如果dis[v] < dis[u] + l[u] [v]，那么更新dis[v]
// 重复上述2，3两个步骤，直到所有可以访问的点都已经访问过
// INF = sys.maxsize
// edges = [[]] # 邻接表存储边
// dis = [] # 记录s到其他点的距离
// visited = {} # 记录访问过的点
// while True:
//     mini = INF
//     u = 0
//     flag = False
//     # 遍历所有未访问过点当中距离最小的
//     for i in range(V):
//         if i not in visited and dis[i] < mini:
//             mini, u = dis[i], i
//             flag = True
//     # 如果没有未访问的点，则退出
//     if not flag:
//         break
//  visited[u] = True
//     for v, l in edges[u]:
//         dis[v] = min(dis[v], dis[u] + l)
// func Dijkstra(G *Graph, dist []int, start int) bool {
// 	dist[start] = 0
// 	visited := make([]bool, G.V)
// 	for {
// 		minD, minU := math.MaxInt32, 0
// 		for i := 0; i < G.V; i++ {
// 			if !visited[i] && dist[i] < minD {
// 				minD = dist[i]
// 				minU = i
// 			}
// 		}
// 		if minD == math.MaxInt32 {
// 			break
// 		}
// 		s := minU
// 		visited[s] = true
// 		for t := 0; t < G.V; t++ {
// 			if !visited[t] && G.Dis[s][t] < math.MaxInt32 {
// 				if G.Dis[s][t] < 0 {
// 					return false
// 				}
// 				if dist[t] > dist[s]+G.Dis[s][t] {
// 					dist[t] = dist[s] + G.Dis[s][t]
// 				}
// 			}
// 		}
// 	}
// 	return true
// }

func floyd(G *Graph, dist [][]int) bool {
	for i := 0; i < G.V; i++ {
		for j := 0; j < G.V; j++ {
			for k := 0; k < G.V; k++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					if i == j && dist[i][j] < 0 {
						return false
					}
				}
			}
		}
	}
	return true
}
