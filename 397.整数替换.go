/*
 * @lc app=leetcode.cn id=397 lang=golang
 *
 * [397] 整数替换
 */

// @lc code=start
func integerReplacement(n int) int {
	memo := make(map[int]int)
	memo[1] = 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var dfs func(num int) int
	dfs = func(num int) int {
		if cnt, ok := memo[num]; ok {
			return cnt
		}
		if n%2 == 0 {
			need := 1 + dfs(num/2)
			memo[num] = need
			return need
		}
		sub := dfs(num - 1)
		plus := dfs(num + 1)
		need := min(plus, sub) + 1
		memo[num] = need
		return need
	}
	ans := dfs(n)
	return ans
}

// @lc code=end

