/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 */

// @lc code=start
func lexicalOrder(n int) []int {
	ans := make([]int, 0)
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur > n {
			return
		}
		ans = append(ans, cur)
		for i := 0; i < 10; i++ {
			dfs(cur*10 + i)
		}
	}
	for i := 1; i < 10; i++ {
		dfs(i)
	}
	return ans
}

// @lc code=end

