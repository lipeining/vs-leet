/*
 * @lc app=leetcode.cn id=1025 lang=golang
 *
 * [1025] 除数博弈
 */

// @lc code=start
func divisorGame(N int) bool {
	dp := make([]bool, N+5)
	// 当前 i 先手的状态，
	dp[1] = false
	dp[2] = true
	for i := 3; i <= N; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !dp[i-j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[N]
	// memo := make(map[int]bool)
	// memo[1] = false
	// var dfs func(num int) bool
	// dfs = func(num int) bool {
	// 	if ret, ok := memo[num]; ok {
	// 		return ret
	// 	}
	// 	for x := 1; x < num; x++ {
	// 		if num%x == 0 {
	// 			next := !dfs(num - x)
	// 			if next {
	// 				return true
	// 			}
	// 		}
	// 	}
	// 	memo[num] = false
	// 	return false
	// }
	// return dfs(N)
}

// @lc code=end

