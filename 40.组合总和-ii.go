/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	if len(candidates) == 0 {
		return ans
	}
	// 不使用集合的话，如何处理呢？
	// 参考题解，使用下标作为筛选条件
	var dfs func(candidates []int, target int, sum int, index int, used []bool, path []int)
	dfs = func(candidates []int, target int, sum int, index int, used []bool, path []int) {
		if sum > target {
			return
		}
		if sum == target {
			fmt.Println(ans, path)
			ans = append(ans, path)
			return
		}
		for i:=index; i < len(candidates); i++ {
			if used[i] {
				continue
			}
			if i>index && candidates[i] == candidates[i-1] {
				continue
			}
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, candidates[i])
			used[i] = true
			dfs(candidates, target, sum + candidates[i], i+1, used, tmp)
			used[i] = false
		}
	}
	sort.Ints(candidates)
	path := make([]int, 0)
	used := make([]bool, len(candidates))
	dfs(candidates, target, 0, 0, used, path)
	return ans
}
// @lc code=end

