/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
    ans := make([][]int, 0)
	if len(candidates) == 0 {
		return ans
	}
	// 不使用集合的话，如何处理呢？
	// 参考题解，使用下标作为筛选条件
	var dfs func(candidates []int, target int, sum int, index int, path []int)
	dfs = func(candidates []int, target int, sum int, index int, path []int) {
		if sum > target {
			return
		}
		if sum == target {
			fmt.Println(ans, path)
			ans = append(ans, path)
			return
		}
		for i:=index; i < len(candidates); i++ {
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, candidates[i])
			dfs(candidates, target, sum + candidates[i], i, tmp)
		}
	}
	path := make([]int, 0)
	dfs(candidates, target, 0, 0, path)
	return ans
}
// @lc code=end

