/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	var dfs func(n int, k int, index int, sum int, path []int)
	dfs = func(n int, k int, index int, sum int, path []int) {
		if len(path) == k {
			if sum == n {
				ans = append(ans, path)
			}
			return
		}
		for i:=index;i<=9;i++ {
			tmp:=make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, i)
			dfs(n,k,i+1,sum+i, tmp)
		}
	}
	path := make([]int, 0)
	dfs(n, k, 1, 0, path)
	return ans
}
// @lc code=end

