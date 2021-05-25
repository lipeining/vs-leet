/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	m := make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && isConnected[i][j] == 1 {
				m[i] = append(m[i], j)
				m[j] = append(m[j], i)
			}
		}
	}
	ans := 0
	seen := make(map[int]bool)
	var dfs func(index int)
	dfs = func(index int) {
		if seen[index] {
			return
		}
		seen[index] = true
		for _, next := range m[index] {
			dfs(next)
		}
	}
	for i := 0; i < n; i++ {
		if !seen[i] {
			ans++
			dfs(i)
		}
	}
	fmt.Println("anssssssssss", ans)
	return ans
}

// @lc code=end

