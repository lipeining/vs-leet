/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 */

// @lc code=start
func getPermutation(n int, k int) string {
	nums := make([]string, 0)
	for i:=1;i<=n;i++ {
		nums = append(nums, strconv.Itoa(i))
	}
	if len(nums) == 0 {
		return ""
	}
	counter := 0
	res := ""
	var dfs func(nums []string, used []bool, path string)
	dfs = func(nums []string, used []bool, path string) {
		if counter >= k {
			return
		}
		if len(path) == len(nums) {
			counter++
			if counter == k {
				res = path
			}
			return
		}
		for i:=0;i<len(nums);i++ {
			if used[i] {
				continue
			}
			used[i] = true
			tmp := path + nums[i]
			dfs(nums, used, tmp)
			used[i] = false
		}
	}
	used := make([]bool, len(nums))
	dfs(nums, used, "")
	return res
}
// @lc code=end

