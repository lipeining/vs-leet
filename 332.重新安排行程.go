/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 */

// @lc code=start
func findItinerary(tickets [][]string) []string {
	graph := getGraph(tickets)
	ans := make([]string, 0)
	var dfs func (graph map[string][]string, start string) 
	dfs = func (graph map[string][]string, start string) {
		list,ok := graph[start]
		if !ok {
			return
		}
		length := len(list)
		for length != 0 {
			head := list[0]
			dfs(graph, head)
			length--
			list = list[1:]
			// graph[start] = graph[start][1:]
		}
		// fmt.Println(graph)
		ans = append([]string{start}, ans...)
	}
	dfs(graph, "JFK")
	// fmt.Println(ans)
	return ans
}

func getGraph(tickets [][]string) map[string][]string {
	m := make(map[string][]string)
	for i:=0;i<len(tickets);i++ {
		from := tickets[i][0]
		to := tickets[i][1]
		_,took := m[to]
		if !took {
			m[to] = make([]string, 0)
		}
		list,ok := m[from]
		if ok {
			m[from] = append(list, to)
		} else {
			list = make([]string, 0)
			list = append(list, to)
			m[from] = list
		}
	}
	for k,_ := range m {
		sort.Strings(m[k])
	}
	fmt.Println(m)
	return m
}
// @lc code=end

