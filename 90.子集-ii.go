/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}
	var dfs func(nums []int, k int, index int, path []int)
	dfs = func(nums []int, k int, index int, path []int) {
		if len(path) == k {
			ans = append(ans, path)
			return
		}
		for i:=index;i<len(nums);i++ {
			if i > index && nums[i-1] == nums[i] {
				continue
			}
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, nums[i])
			dfs(nums, k, i+1, tmp)
		}
	}
	path := make([]int, 0)
	sort.Ints(nums)
	for i:=0;i<=len(nums);i++ {
		dfs(nums, i, 0, path)
	}
	return ans
}
// @lc code=end

