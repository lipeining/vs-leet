/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	nums := make([]int, 0)
	for i:=1;i<=n;i++ {
		nums = append(nums, i)
	}
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
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, nums[i])
			dfs(nums, k, i+1, tmp)
		}
	}
	path := make([]int, 0)
	dfs(nums, k, 0, path)
	return ans
	// var dfs func(nums []int, k int, index int, used []bool, path []int)
	// dfs = func(nums []int, k int, index int,used []bool, path []int) {
	// 	if len(path) == k {
	// 		ans = append(ans, path)
	// 		return
	// 	}
	// 	for i:=index;i<len(nums);i++ {
	// 		tmp := make([]int, len(path))
	// 		copy(tmp, path)
	// 		tmp = append(tmp, nums[i])
	// 		used[i]=true
	// 		dfs(nums, k, i+1, used, tmp)
	// 		used[i]=false
	// 	}
	// }
	// used := make([]bool, len(nums))
	// path := make([]int, 0)
	// dfs(nums, k, 0, used, path)
	// return ans
}
// @lc code=end

