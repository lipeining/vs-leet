/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	// ans := false
	// 超时了，最后一个测试用例
	// var dfs func(nums []int, now int)
	// dfs = func(nums []int, now int) {
	// 	if ans {
	// 		return
	// 	}
	// 	if now >= len(nums) - 1 {
	// 		ans = true
	// 		return
	// 	}
	// 	step := nums[now]
	// 	for i:=step;i>=1;i-- {
	// 		dfs(nums, now+i)
	// 	}
	// }
	// dfs(nums, 0)
	// return ans
	// return dfs(nums, 0)
	
	
	// dp
	if len(nums) <= 1 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	for i:=0;i<len(nums);i++ {
		if !dp[i] {
			return false
		}
		step := nums[i]
		if i+step >= len(nums) {
			step = len(nums)-i-1
		}
		for j:=1;j<=step;j++ {
			dp[i+j] = true
		}
	}
	return dp[len(nums)-1]
}
func dfs(nums []int, now int) bool {
	if now >= len(nums)-1 {
		return true
	}
	max := nums[now] + now
	if max > len(nums) {
		max = len(nums)
	}
	for i:=now+1;i<=max;i++ {
		if dfs(nums, i) {
			return true
		}
	}
	return false
}
// @lc code=end

