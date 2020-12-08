/*
 * @lc app=leetcode.cn id=464 lang=golang
 *
 * [464] 我能赢吗
 */
package main

import (
	"strings"
)

// func main() {
// 	fmt.Println(canIWin(10, 11))
// 	fmt.Println(canIWin(10, 3))
// }

// @lc code=start
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	memo := make(map[string]bool)
	var dfs func(used []string, d int) bool
	dfs = func(used []string, d int) bool {
		key := strings.Join(used, ",")
		// fmt.Println(memo)
		if res, ok := memo[key]; ok {
			return res
		}
		for i := 1; i <= maxChoosableInteger; i++ {
			if used[i] == "1" {
				continue
			}
			used[i] = "1"
			if d <= i || !dfs(used, d-i) {
				memo[key] = true
				used[i] = ""
				return true
			}
			used[i] = ""
		}
		memo[key] = false
		return false
	}
	used := make([]string, maxChoosableInteger+1)
	return dfs(used, desiredTotal)
}

// @lc code=end
