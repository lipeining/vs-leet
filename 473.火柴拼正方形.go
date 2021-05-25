/*
 * @lc app=leetcode.cn id=473 lang=golang
 *
 * [473] 火柴拼正方形
 */

// @lc code=start
func makesquare(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%4 != 0 {
		return false
	}
	part := sum / 4
	edge := make([]int, 4)
	check := func() bool {
		for i := 0; i < 4; i++ {
			if edge[i] != part {
				return false
			}
		}
		return true
	}
	ans := false
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index == n {
			return check()
		}
		length := nums[index]
		for i := 0; i < 4; i++ {
			if edge[i]+length <= part {
				edge[i] += length
				if dfs(index + 1) {
					return true
				}
				edge[i] -= length
			}
		}
		return false
	}
	ans = dfs(0)
	return ans
}

// @lc code=end

