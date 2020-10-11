package template

// 121 122 123 188 309 714 股票问题
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-l-3/
// base case：
// dp[-1][k][0] = dp[i][0][0] = 0
// dp[-1][k][1] = dp[i][0][1] = -infinity

// 状态转移方程：
// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

// k 无次数限制时，对于 dp[0][j][1] = math.MinInt32
// 即，i,j为0时，需要为最小值
// for i:=0;i<=n;i++ {
// 	dp[i] = make([][]int, k+1)
// 	for j:=0;j<=k;j++ {
// 			dp[i][j] = make([]int, 2)
// 			dp[0][j][1] = math.MinInt32
// 	}
// 	// dp[i][0][0] = 0
// 	dp[i][0][1] = math.MinInt32
// }

// 作者：labuladong
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-l-3/

// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// 121 122 123 188 309 714 股票问题
