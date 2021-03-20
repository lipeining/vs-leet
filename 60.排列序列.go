/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 */

// @lc code=start
func getPermutation(n int, k int) string {
	fact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = i * fact[i-1]
	}
	used := make([]bool, n+1)
	var dfs func(index int, path string) string
	dfs = func(index int, path string) string {
		if index == n {
			return path
		}
		cnt := fact[n-1-index]
		for i := 1; i <= n; i++ {
			if used[i] {
				continue
			}
			if cnt < k {
				k -= cnt
				continue
			}
			path += strconv.Itoa(i)
			used[i] = true
			return dfs(index+1, path)
		}
		return ""
	}
	return dfs(0, "")
}

// @lc code=end

