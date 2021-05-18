package main

import "math"

/*
 * @lc app=leetcode.cn id=1000 lang=golang
 *
 * [1000] 合并石头的最低成本
 */

// @lc code=start
func mergeStones(stones []int, K int) int {
	length := len(stones)
	if (length-1)%(K-1) != 0 {
		return -1
	}
	sums := make([]int, length+1)
	for i := 1; i <= length; i++ {
		sums[i] = sums[i-1] + stones[i-1]
	}
	dp := make([][][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([][]int, length)
		for j := 0; j < length; j++ {
			dp[i][j] = make([]int, K+1)
			for t := 0; t <= K; t++ {
				dp[i][j][t] = math.MaxInt32
			}
		}
	}
	for i := 0; i < length; i++ {
		dp[i][i][1] = 0
		// j := K - 1
		// if j < length {
		// 	dp[i][i+j][1] = sums[i+j] - sums[i]
		// }
	}
	for i := 2; i <= length; i++ {
		j := i + length - 1
		for k := 2; k <= K; k++ {
			for m := i; m < j; m += K - 1 {
				dp[i][j][k] = min(dp[i][j][k], dp[i][m][1]+dp[m+1][j][k-1])
				dp[i][j][1] = sums[j] - sums[i] + dp[i][j][k]
			}
		}
	}
	return dp[0][length-1][1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
// for(int len =2;len <= N; len++){ // sub problem length
// 	for(int i=0;i<=N-len;i++){
// 		int j = i+len-1;
// 		for(int k=2;k<=K;k++){
// 			for(int m=i;m<j;m +=K-1){ // m 跳步应该是k-1
// 				dp[i][j][k]=Math.min(dp[i][j][k],dp[i][m][1]+dp[m+1][j][k-1]);
// 				dp[i][j][1]=sum(i,j,preSum)+dp[i][j][k];
// 			}
// 		}
// 	}
// }
// return dp[0][N-1][1];

// 作者：xiaobinbin
// 链接：https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/solution/jing-dian-de-shi-tou-you-xi-qu-jian-dpdan-bu-shi-y/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// 而要分成k分 就从i到j之间找一个t 把i到t分成k-1份 再把t+1到j合成一份,递归计算两边要合并的值为多少，再加起来
// 所以我们考虑每个t 看哪一个t使得两边合并的成本再加起来是最小的 就取这个t
// 也就是：
// dp(i,j,k)=min{dp(i,t,k-1)+dp(t+1,j,1)|t from i to j-1}

// 合并一下用递推表达式表示就是：
// dp(i,j,k)=min{dp(i,t,k-1)+dp(t+1,j,1)|t from i to j-1}
// dp(i,j,1)=dp(i,j,k)+sum(i,j)

// 初始条件是
// dp(i,i,1)=0 只有一个的时候是不需要合并的
// dp(i,i+k-1,1)=sum(i,i+k-1) 这个很显然

// 作者：JiaQiaoY
// 链接：https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/solution/dp-dp-dppppp-by-jiaqiaoy/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
