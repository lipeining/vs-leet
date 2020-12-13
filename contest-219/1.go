package main

import "fmt"

func main() {
	fmt.Println("ans")
	// 	输入：stones = [5,3,1,4,2]
	// 输出：6
	// 解释：
	// - 爱丽丝移除 2 ，得分 5 + 3 + 1 + 4 = 13 。游戏情况：爱丽丝 = 13 ，鲍勃 = 0 ，石子 = [5,3,1,4] 。
	// - 鲍勃移除 5 ，得分 3 + 1 + 4 = 8 。游戏情况：爱丽丝 = 13 ，鲍勃 = 8 ，石子 = [3,1,4] 。
	// - 爱丽丝移除 3 ，得分 1 + 4 = 5 。游戏情况：爱丽丝 = 18 ，鲍勃 = 8 ，石子 = [1,4] 。
	// - 鲍勃移除 1 ，得分 4 。游戏情况：爱丽丝 = 18 ，鲍勃 = 12 ，石子 = [4] 。
	// - 爱丽丝移除 4 ，得分 0 。游戏情况：爱丽丝 = 18 ，鲍勃 = 12 ，石子 = [] 。
	// 得分的差值 18 - 12 = 6 。
	// 输入：stones = [7,90,5,1,100,10,10,2]
	// 输出：122
	fmt.Println("ans::", stoneGameVII([]int{5, 3, 1, 4, 2}))
	fmt.Println("ans::", stoneGameVII([]int{7, 90, 5, 1, 100, 10, 10, 2}))
}
func stoneGameVII(stones []int) int {
	n := len(stones)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + stones[i]
	}
	// fmt.Println(pre)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		// 只剩一个时，价值为 0
		// dp[i][i] = stones[i]
	}
	// \textit{dp}[i][j]dp[i][j]
	// 表示当剩下的石子堆为下标 ii 到下标 jj 时，当前玩家与另一个玩家的价值之差的最大值
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			iValue := pre[j+1] - pre[i+1]
			jValue := pre[j] - pre[i]
			// fmt.Println("i:", i, "iValue:", iValue, "j:", j, "jValue:", jValue)
			dp[i][j] = max(iValue-dp[i+1][j], jValue-dp[i][j-1])
		}
	}
	// fmt.Println(dp)
	return dp[0][n-1]
}
