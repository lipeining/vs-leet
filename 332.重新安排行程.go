/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 */

// @lc code=start
func findItinerary(tickets [][]string) []string {
	graph := getGraph(tickets)
	ans := make([]string, 0)
	var dfs func(start string)
	dfs = func(start string) {
		for {
			if v, ok := graph[start]; !ok || len(v) == 0 {
				break
			}
			tmp := graph[start][0]
			graph[start] = graph[start][1:]
			dfs(tmp)
		}
		// fmt.Println(graph)
		ans = append([]string{start}, ans...)
	}
	dfs("JFK")
	// fmt.Println(ans)
	return ans
}

func getGraph(tickets [][]string) map[string][]string {
	m := make(map[string][]string)
	for i := 0; i < len(tickets); i++ {
		from := tickets[i][0]
		to := tickets[i][1]
		m[from] = append(m[from], to)
	}
	for k, _ := range m {
		sort.Strings(m[k])
	}
	// fmt.Println(m)
	return m
}

// @lc code=end

