/*
 * @lc app=leetcode.cn id=488 lang=golang
 *
 * [488] 祖玛游戏
 */
package main

import "fmt"

// func main() {
// 	findMinStep("WRRBBW", "RB")
// 	findMinStep("WWRRBBWW", "WRBRW")
// 	// findMinStep("G", "GGGGG")
// 	// findMinStep("RBYYBBRRB", "YRBGB")
// }

// @lc code=start
func findMinStep(board string, hand string) int {
	memo := make(map[string]map[string]int)
	// memo := make([][][][]int, len(board)) 思维memo
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var earse func(str string) string
	earse = func(str string) string {
		// 递归处理重复三个元素以上的
		size := len(str)
		i := 0
		trans := ""
		for i < size {
			// if i+1 < size && str[i+1] == str[i] {
			// 	if i+2 < size && str[i+2] == str[i] {
			// 		i = i + 3
			// 		continue
			// 	}
			// }
			// trans += string(str[i])
			i++
		}
		// fmt.Println("earse", str, trans)
		if trans == str {
			return str
		}
		return earse(trans)
	}
	var dfs func(b, h string) int
	dfs = func(b, h string) int {
		// if memo[b] == nil {
		// 	memo[b] = make(map[string]int)
		// }
		if bm, ok := memo[b]; ok {
			if v, ok2 := bm[h]; ok2 {
				return v
			}
		} else {
			memo[b] = make(map[string]int)
		}
		fmt.Println("for b, h start", b, h)
		if len(h) == 0 {
			memo[b][h] = -1
			return -1
		}
		// 求最小
		now := -1
		for i := 0; i < len(h); i++ {
			// h任意球 到 b任意位置
			nh := h[0:i] + h[i+1:]
			char := string(h[i])
			for j := 0; j <= len(b); j++ {
				nb := b[0:j] + char + b[j:]
				nbb := earse(nb)
				if ret := dfs(nbb, nh); ret != -1 {
					fmt.Println("for b, h process", nbb, nh, ret)
					return -1
					if now == -1 {
						now = ret
					} else {
						now = min(now, ret)
					}
				}
			}
			return -1
		}
		memo[b][h] = now
		fmt.Println("for b, h end", b, h, now)
		return now
	}
	ans := dfs(board, hand)
	fmt.Println(memo)
	fmt.Println("ansssss", ans)
	return ans
}

// @lc code=end
