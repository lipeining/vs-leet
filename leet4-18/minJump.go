package main

import (
	"fmt"
	"math"
)

func main() {
	// [2, 5, 1, 1, 1, 1] 3
	jump := []int{2, 5, 1, 1, 1, 1}
	ans := minJump(jump)
	fmt.Println(ans)
}
func minJump(jump []int) int {
	// 从右到左的动态规划，以次数为最终结果进行合理计算
	// minJump
	// 当前节点为 i ，可以选择的是跳到 i + jumps[i] 点
	end := false
	dp := make([]int, len(jump))
	for i := 0; i < len(jump); i++ {
		dp[i] = math.MaxInt32
	}
	// 到达 0 不需要步骤
	dp[0] = 0
	N := len(jump) + 1
	ans := math.MaxInt32
	for !end {
		for i := 0; i < len(jump); i++ {
			for j := 0; j < len(jump); j++ {
				if j >= i {
					dp[i] = min(dp[i], dp[j]+1)
				} else if i == jump[j]+j {
					dp[i] = min(dp[i], dp[j]+1)
				} else {
					// i 与 j 无关系
				}
			}
		}
		// fmt.Println(dp)
		all := true
		for i := 0; i < len(jump); i++ {
			if dp[i] == math.MaxInt32 {
				all = false
			} else if dp[i]+jump[i] >= N {
				ans = min(ans, dp[i]+1)
			}
		}
		if ans != math.MaxInt32 && all {
			end = true
		}
	}
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
