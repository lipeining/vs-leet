/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}
	var dfs func(nums []int, used []bool, path []int)
	dfs = func(nums []int, used []bool, path []int) {
		if len(path) == len(nums) {
			ans = append(ans, path)
			return
		}
		for i:=0;i<len(nums);i++ {
			if used[i] {
				continue
			}
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, nums[i])
			used[i] = true
			dfs(nums, used, tmp)
			used[i] = false
		}
	}
	path := make([]int, 0)
	used := make([]bool, len(nums))
	dfs(nums, used, path)
	return ans
}
// @lc code=end

