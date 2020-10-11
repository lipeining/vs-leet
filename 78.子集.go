/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}
	var dfs func(nums []int, index int, path []int)
	dfs = func(nums []int, index int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)
		for i:=index;i<len(nums);i++ {
			path = append(path, nums[i])
			dfs(nums, i+1, path)
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	dfs(nums, 0, path)
	return ans
}
// func subsets(nums []int) [][]int {
// 	// nums := make([]int, 0)
// 	// for i:=1;i<=n;i++ {
// 	// 	nums = append(nums, i)
// 	// }
// 	ans := make([][]int, 0)
// 	if len(nums) == 0 {
// 		return ans
// 	}
// 	var dfs func(nums []int, k int, index int, path []int)
// 	dfs = func(nums []int, k int, index int, path []int) {
// 		if len(path) == k {
// 			ans = append(ans, path)
// 			return
// 		}
// 		for i:=index;i<len(nums);i++ {
// 			tmp := make([]int, len(path))
// 			copy(tmp, path)
// 			tmp = append(tmp, nums[i])
// 			dfs(nums, k, i+1, tmp)
// 		}
// 	}
// 	path := make([]int, 0)
// 	for i:=0;i<=len(nums);i++ {
// 		dfs(nums, i, 0, path)
// 	}
// 	return ans
// }
// @lc code=end

