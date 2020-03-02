/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

// @lc code=start
func maxCoins(nums []int) int {
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for len := 2; len < n; len++ {
		for i := 0; i < n-len; i++ {
			j := i + len
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[j]*nums[k])
			}
		}
	}
	return dp[0][n-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

// dp[i][j] 表示戳破 [i+1...j-1] 号气球的最大收益。

// 假设 k 号气球（i+1 <= k <= j-1）是 [i+1...j-1] 中最后一个被戳破的，则

// dp[i][j] = max {for k = range(i+1, j -1) nums[i] * nums[k] * nums[j] + dp[i][k] + dp[k][j]}

// class Solution {
// public:
//     int maxCoins(vector<int>& nums) {
//         nums.insert(nums.begin(), 1);
//         nums.push_back(1);

//         int n = nums.size();
//         vector<vector<int>> dp(n, vector<int>(n));

//         // 这里不能按行打表
//         // 只能按长度打表, dp[i][j], j-i >= 2
//         for (int len = 2; len < n; ++len) {
//             for (int i = 0; i < n - len; ++i) {
//                 int j = i + len;
//                 for (int k = i + 1; k < j; ++k) {
//                     dp[i][j] = max(dp[i][j], nums[i] * nums[k] * nums[j] + dp[i][k] + dp[k][j]);
//                 }
//             }
//         }
//         return dp[0][n-1];
//     }
// };

// 作者：ivan_allen
// 链接：https://leetcode-cn.com/problems/burst-balloons/solution/dong-tai-gui-hua-by-ivan_allen-2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。