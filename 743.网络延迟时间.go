/*
 * @lc app=leetcode.cn id=743 lang=golang
 *
 * [743] 网络延迟时间
 */

// @lc code=start
func networkDelayTime(times [][]int, N int, K int) int {
	// graph[i]保存i节点的相邻节点和距离元组的列表
	graph := make(map[int][][]int)
	for _, time := range times {
		graph[time[0]] = append(graph[time[0]], []int{time[1], time[2]})
	}
	dist := make(map[int]int)
	minHeap := &NodeDistMinHeap{}
	heap.Push(minHeap, NodeDist{
		node: K,
		dist: 0,
	})
	for minHeap.Len() > 0 {
		nd := heap.Pop(minHeap).(NodeDist)
		if _, ok := dist[nd.node]; ok {
			continue
		}
		dist[nd.node] = nd.dist
		for _, tuple := range graph[nd.node] {
			if _, ok := dist[tuple[0]]; !ok {
				heap.Push(minHeap, NodeDist{
					node: tuple[0],
					dist: tuple[1] + nd.dist,
				})
			}
		}
	}
	if len(dist) != N {
		return -1
	}
	max := -1
	for _, d := range dist {
		if d > max {
			max = d
		}
	}
	return max
}

type NodeDist struct {
	node int
	dist int
}

type NodeDistMinHeap []NodeDist

func (pq *NodeDistMinHeap) Len() int {
	return len(*pq)
}
func (pq *NodeDistMinHeap) Less(i, j int) bool {
	return (*pq)[i].dist < (*pq)[j].dist
}
func (pq *NodeDistMinHeap) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *NodeDistMinHeap) Push(x interface{}) {
	*pq = append(*pq, x.(NodeDist))
}

func (pq *NodeDistMinHeap) Pop() interface{} {
	n := len(*pq) - 1
	x := (*pq)[n]
	*pq = (*pq)[:n]
	return x
}
func (pq *NodeDistMinHeap) Peek() NodeDist {
	return (*pq)[0]
}

// 作者：hundanLi
// 链接：https://leetcode-cn.com/problems/network-delay-time/solution/wang-luo-yan-chi-shi-jian-by-hundanli/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
// @lc code=end
// func networkDelayTime(times [][]int, N int, K int) int {
	
// 	// dfs 返回的是需要的时间，如果为 -1 表明无法传播出去了
// 	// 然后取最大值作为该节点的时间值
// 	ans := 0
// 	var dfs func(times [][]int, used []bool, visited []bool, N int, now int, index int) int
// 	dfs = func(times [][]int, used []bool, visited []bool, N int, now int, index int) int {
// 		time := times[index]
// 		start := time[1]
// 		fmt.Println(now, N)
// 		if now == N {
// 			// 检查 visited 是否是一条完整的路线
// 			flag := true
// 			for i:=1;i<len(visited);i++ {
// 				if !visited[i] {
// 					flag = false
// 					break
// 				}
// 			}
// 			if !flag {
// 				return -1
// 			}
// 			return time[2]
// 		}
// 		i:=0
// 		max := -1
// 		hasPath := false
// 		for i<len(times) {
// 			if used[i] {
// 				continue
// 			}
// 			if times[i][0] != start {
// 				continue
// 			}
// 			hasPath = true
// 			used[i] = true
// 			visited[start] = true
// 			tmp := dfs(times, used, visited, N, now + 1, i)
// 			used[i] = false
// 			visited[start] = false
// 			i++
// 			if tmp != -1 && tmp > max {
// 				max = tmp
// 			}
// 		}
// 		if !hasPath {
// 			return -1
// 		}
// 		if max == -1 {
// 			return max
// 		}
// 		return max + time[2]
// 	}
// 	used := make([]bool, len(times))
// 	visited := make([]bool, N + 1)
// 	for i:=0;i<len(times);i++ {
// 		if times[i][0] == K {
// 			time := dfs(times, used, visited, N, 1, i)
// 			if time < ans {
// 				ans = time
// 			}
// 		}
// 	}	
// 	return ans
// }
