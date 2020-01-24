/*
 * @lc app=leetcode.cn id=526 lang=golang
 *
 * [526] 优美的排列
 */

// @lc code=start
func countArrangement(N int) int {
	ans := 0
	var dfs func(N int, now int, used []bool)
	dfs = func(N int, now int, used []bool) {
		if now == N {
			ans++
			return
		}
		for i:=1;i<=N;i++ {
			if used[i] {
				continue
			}
			index := now + 1
			if check(index, i) {
				used[i] = true
				dfs(N, now+1, used)
				used[i] = false
			}
		}
	}
	used := make([]bool, N+1)
	dfs(N, 0, used)
	return ans
}
func check(a int, b int) bool {
	return a%b==0 || b%a==0
}
// @lc code=end

