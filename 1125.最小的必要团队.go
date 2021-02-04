/*
 * @lc app=leetcode.cn id=1125 lang=golang
 *
 * [1125] 最小的必要团队
 */
package main

// func main() {
// 	smallestSufficientTeam([]string{"java", "nodejs", "reactjs"},
// 		[][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}})
// }

// @lc code=start
func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	m := len(req_skills)
	// min := func(a, b int) int {
	// 	if a < b {
	// 		return a
	// 	}
	// 	return b
	// }
	rmap := make(map[string]int)
	for index, skill := range req_skills {
		rmap[skill] = index
	}
	n := 1 << m
	dp := make([]int, n)
	path := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
		path[i] = make([]int, 0)
	}
	path[0] = make([]int, 0)
	dp[0] = 0
	for i := 0; i < m; i++ {
		idx := 0
		for _, s := range people[i] {
			if index, ok := rmap[s]; ok {
				idx |= 1 << index
			}
		}
		for j := 0; j < n; j++ {
			if dp[j] >= 0 {
				x := j | idx
				// 团队合并
				if dp[x] == -1 || dp[x] > dp[j]+1 {
					dp[x] = dp[j] + 1
					path[x] = make([]int, len(path[j]))
					copy(path[x], path[j])
					path[x] = append(path[x], i)
				}
			}
		}
	}
	// fmt.Println(dp)
	// fmt.Println(dp[n-1], path[n-1])
	return path[n-1]
}

// @lc code=end
