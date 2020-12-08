/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */
package main

import "fmt"

// func main() {
// 	findMaxForm([]string{"10", "0", "1"}, 1, 1)
// }

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	// 背包问题啊 0 1
	// length := len(strs)
	// 对于第 i 个选择，拼接或者不拼接
	// strs 相当于 weight
	list := make([][]int, len(strs))
	for i := 0; i < len(strs); i++ {
		z, o := 0, 0
		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '0' {
				z++
			} else {
				o++
			}
		}
		list[i] = []int{z, o}
	}
	// fmt.Println(list)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// 以 m , n 最多可以拼接出的字符串数量
	// 对比了官方思路，应该是错在 i,j 应该从大到小排序，而不是从小到大
	//  同时，应该 strs 放在外层
	for k := 0; k < len(strs); k++ {
		z, o := list[k][0], list[k][1]
		for i := m; i >= 0; i-- {
			for j := n; j >= 0; j-- {
				if i >= z && j >= o {
					dp[i][j] = max(dp[i][j], dp[i-z][j-o]+1)
				} else {
					// dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				}
			}
		}
	}

	// 如果需要参考前一层的话，需要三维数组
	// 	for (int i = 1; i <= len; i++) {
	// 		// 注意：有一位偏移
	// 		int[] cnt = countZeroAndOne(strs[i - 1]);
	// 		for (int j = 0; j <= m; j++) {
	// 			for (int k = 0; k <= n; k++) {
	// 				// 先把上一行抄下来
	// 				dp[i][j][k] = dp[i - 1][j][k];

	// 				int zeros = cnt[0];
	// 				int ones = cnt[1];

	// 				if (j >= zeros && k >= ones) {
	// 					dp[i][j][k] = Math.max(dp[i - 1][j][k], dp[i - 1][j - zeros][k - ones] + 1);
	// 				}
	// 			}
	// 		}
	// 	}
	// 	return dp[len][m][n];

	// 作者：liweiwei1419
	// 链接：https://leetcode-cn.com/problems/ones-and-zeroes/solution/dong-tai-gui-hua-zhuan-huan-wei-0-1-bei-bao-wen-ti/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	// for i := 1; i <= m; i++ {
	// 	for j := 1; j <= n; j++ {
	// 		for k := 0; k < len(strs); k++ {
	// 			z, o := list[k][0], list[k][1]
	// 			if i >= z && j >= o {
	// 				dp[i][j] = max(dp[i][j], dp[i-z][j-o]+1)
	// 			} else {
	// 				// dp[i][j] = max(dp[i-1][j], dp[i][j-1])
	// 			}
	// 		}
	// 	}
	// }
	fmt.Println(dp)
	return dp[m][n]
	// dp := make([]int, m+n)
	// for i := 1; i <= m+n; i++ {
	// 	for j := 0; j < len(strs); j++ {
	// 		z, o := list[j][0], list[j][1]
	// 	}
	// }
}
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end
