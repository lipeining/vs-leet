/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	ans := make([][]int, 0)
	var dfs func(nums []int, index int, used []bool, path []int)
	dfs = func(nums []int, index int, used []bool, path []int) {
		length := len(path)
		if length >= 2 {
			ans = append(ans, path)
		}
		for i:=index;i<len(nums);i++ {
			if i > index && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			last := nums[i]
			if length > 0 {
				last = path[length-1]
			}
			if nums[i] >= last {
				tmp := make([]int, length)
				copy(tmp, path)
				tmp = append(tmp, nums[i])
				used[i] = true
				dfs(nums, i+1, used, tmp)
				used[i] = false
			}
		}
	}
	// [1,2,3,4,5,6,7,8,9,10,1,1,1,1,1]
	path := make([]int, 0)
	used := make([]bool, len(nums))
	dfs(nums, 0, used, path)
	return ans
}
// @lc code=end

