/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
    var dfs func(index int, path []int)
    dfs = func(index int, path []int) {
        c := make([]int, len(path))
        copy(c, path)
        ans = append(ans, c)
        for j:=index;j<len(nums);j++ {
            path = append(path, nums[j])
            dfs(j+1, path)
            path = path[:len(path)-1]
        }
    }
    path := make([]int, 0)
    dfs(0, path)
    return ans
	// // nums := make([]int, 0)
	// // for i:=1;i<=n;i++ {
	// // 	nums = append(nums, i)
	// // }
	// ans := make([][]int, 0)
	// if len(nums) == 0 {
	// 	return ans
	// }
	// var dfs func(nums []int, k int, index int, path []int)
	// dfs = func(nums []int, k int, index int, path []int) {
	// 	if len(path) == k {
	// 		ans = append(ans, path)
	// 		return
	// 	}
	// 	for i:=index;i<len(nums);i++ {
	// 		tmp := make([]int, len(path))
	// 		copy(tmp, path)
	// 		tmp = append(tmp, nums[i])
	// 		dfs(nums, k, i+1, tmp)
	// 	}
	// }
	// path := make([]int, 0)
	// for i:=0;i<=len(nums);i++ {
	// 	dfs(nums, i, 0, path)
	// }
	// return ans
}
// @lc code=end

