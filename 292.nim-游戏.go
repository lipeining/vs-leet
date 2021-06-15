/*
 * @lc app=leetcode.cn id=292 lang=golang
 *
 * [292] Nim 游戏
 */

// @lc code=start
// func canWinNim(n int) bool {
// 	memo := make(map[int]bool)
// 	var dfs func(num int)bool 
// 	dfs = func(num int)bool {
// 		if ret,ok := memo[num];ok {
// 			return ret
// 		}
// 		if num <= 3 {
// 			memo[num] = true
// 			return true
// 		}
// 		for i := 1; i <= 3; i++ {
// 			if !canWinNim(num - i) {
// 				memo[num-i] = true
// 				return true
// 			}
// 		}
// 		memo[num] = false
// 		return false
// 	}
// 	return dfs(n)
// }

func canWinNim(n int) bool {
    return n % 4 != 0
}
// @lc code=end

