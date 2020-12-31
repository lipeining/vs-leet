/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return ans
	}
	var dfs func(nums []int, used []bool, path []int)
	dfs = func(nums []int, used []bool, path []int) {
		if len(path) == len(nums) {
			c := make([]int, len(path))
			copy(c, path)
			ans = append(ans, c)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			dfs(nums, used, path)
			used[i] = false
			path = path[:len(path)-1]
		}
		// if len(path) == len(nums) {
		// 	ans = append(ans, path)
		// 	return
		// }
		// for i:=0;i<len(nums);i++ {
		// 	if used[i] {
		// 		continue
		// 	}
		// 	if i > 0 && nums[i-1]==nums[i] && !used[i-1]{
		// 		continue
		// 	}
		// 	tmp := make([]int, len(path))
		// 	copy(tmp, path)
		// 	tmp = append(tmp,  nums[i])
		// 	used[i]=true
		// 	dfs(nums, used, tmp)
		// 	used[i]=false
		// }
	}
	sort.Ints(nums)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	dfs(nums, used, path)
	return ans
}

// @lc code=end

