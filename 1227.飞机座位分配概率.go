/*
 * @lc app=leetcode.cn id=1227 lang=golang
 *
 * [1227] 飞机座位分配概率
 */
package main

// import "fmt"

// func main() {
// 	nthPersonGetsNthSeat(1)
// 	nthPersonGetsNthSeat(2)
// 	nthPersonGetsNthSeat(3)
// 	nthPersonGetsNthSeat(4)
// 	nthPersonGetsNthSeat(10)
// }

// @lc code=start
func nthPersonGetsNthSeat(n int) float64 {
	if n >= 2 {
		return 0.5
	}
	return 1
	// p := make([]float64, 100001)
	// p[1] = 1.0
	// p[2] = 0.5
	// for i := 3; i <= n; i++ {
	// 	var sum float64 = 1.0 / float64(i)
	// 	for k := 2; k < i; k++ {
	// 		sum += p[i-k+1] / float64(i)
	// 	}
	// 	p[i] = sum
	// }
	// return p[n]
	// dp := make([][]float64, n+1)
	// for i := 0; i <= n; i++ {
	// 	dp[i] = make([]float64, 2)
	// }
	// // dp[i][0] 坐对的概率 dp[i][1] 坐错的概率
	// dp[1][0] = float64(1) / float64(n)
	// dp[1][1] = float64(1) - dp[1][0]
	// for i := 2; i <= n; i++ {
	// 	dp[i][1] = dp[i-1][1] / float64(n-i+1)
	// 	dp[i][0] = float64(1) - dp[i][1]
	// 	// fmt.Println("on i", i, dp[i])
	// }
	// // fmt.Println("ansssss", dp)
	// return dp[n][0]
}

// @lc code=end
