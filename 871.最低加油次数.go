/*
 * @lc app=leetcode.cn id=871 lang=golang
 *
 * [871] 最低加油次数
 */

// @lc code=start
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	// dp[i] 为加 i 次油能走的最远距离，
	// 需要满足 dp[i] >= target 的最小 i。
	length := len(stations)
	dp := make([]int, length+1)
	dp[0] = startFuel
	for i := 0; i < length; i++ {
		for t := i; t >= 0; t-- {
			// 加油能够到达这个地点 0 先
			if dp[t] >= stations[i][0] {
				dp[t+1] = max(dp[t+1], dp[t]+stations[i][1])
			}
		}
	}
	for i := 0; i <= length; i++ {
		if dp[i] >= target {
			return i
		}
	}
	return -1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
// class Solution {
//     public int minRefuelStops(int target, int startFuel, int[][] stations) {
//         int N = stations.length;
//         long[] dp = new long[N + 1];
//         dp[0] = startFuel;
//         for (int i = 0; i < N; ++i)
//             for (int t = i; t >= 0; --t)
//                 if (dp[t] >= stations[i][0])
//                     dp[t+1] = Math.max(dp[t+1], dp[t] + (long) stations[i][1]);

//         for (int i = 0; i <= N; ++i)
//             if (dp[i] >= target) return i;
//         return -1;
//     }
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/minimum-number-of-refueling-stops/solution/zui-di-jia-you-ci-shu-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
