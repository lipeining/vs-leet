/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(index, last int)
	dfs = func(index, last int) {
		if index == len(nums) {
			if len(path) >= 2 {
				t := make([]int, len(path))
				copy(t, path)
				ans = append(ans, t)
			}
			return
		}
		if nums[index] >= last {
			path = append(path, nums[index])
			dfs(index+1, nums[index])
			path = path[:len(path)-1]
		}
		if nums[index] != last {
			dfs(index+1, last)
		}
	}
	// [1,2,3,4,5,6,7,8,9,10,1,1,1,1,1]
	dfs(0, math.MinInt32)
	return ans
}
// @lc code=end

