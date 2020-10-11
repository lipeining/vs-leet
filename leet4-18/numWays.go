func numWays(n int, relation [][]int, k int) int {
	ans := 0
	var dfs func(id, now int)
	dfs = func(id, now int) {
		if now == k {
			if id == n-1 {
				ans++
			}
			return
		}
		for _,r := range relation {
			if r[0] == id {
				// 必须当前节点开头
				dfs(r[1], now+1)
			}
		}
	}
	dfs(0, 0)
	return ans
}