/*
 * @lc app=leetcode.cn id=799 lang=golang
 *
 * [799] 香槟塔
 */

// @lc code=start
func champagneTower(poured int, query_row int, query_glass int) float64 {
	if poured == 0 {
		return 0
	}
	n := 100
	dp := make([]float64, n)
	dp[0] = float64(poured)
	for i := 0; i < query_row; i++ {
		newDP := make([]float64, n)
		for j := 0; j <= i; j++ {
			if dp[j] < 1 {
				continue
			} else {
				node := dp[j]
				half := (node - 1) / 2
				newDP[j] += half
				newDP[j+1] += half
			}
		}
		dp = newDP
		// fmt.Println(dp)
	}
	ans := dp[query_glass]
	if ans >= 1 {
		return 1
	}
	// fmt.Println(ans)
	return ans
	// if poured == 0 {
	// 	return 0
	// }
	// n := 100
	// dp := make([][]float64, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([]float64, n)
	// }
	// dp[0][0] = float64(poured)
	// for i := 0; i < n; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		if dp[i][j] < 1 {
	// 			continue
	// 		} else {
	// 			node := dp[i][j]
	// 			half := (node - 1) / 2
	// 			dp[i+1][j] += half
	// 			dp[i+1][j+1] += half
	// 			dp[i][j] = 1
	// 		}
	// 	}
	// 	fmt.Println(dp[i])
	// }
	// ans := dp[query_row][query_glass]
	// fmt.Println(ans)
	// return ans
}

// @lc code=end

