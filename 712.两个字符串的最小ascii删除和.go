/*
 * @lc app=leetcode.cn id=712 lang=golang
 *
 * [712] 两个字符串的最小ASCII删除和
 */

// @lc code=start
func minimumDeleteSum(s1 string, s2 string) int {
	// fmt.Println(s1[0], int(s1[0])) // 155 155
	// fmt.Println('s', int('s'))
	m := len(s1)
	n := len(s2)
	if m == 0 {
		return getASCII(s2)
	}
	if n == 0 {
		return getASCII(s1)
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := m - 1; i >= 0; i-- {
		dp[i][n] = dp[i+1][n] + getASCII(string(s1[i]))
	}
	for i := n - 1; i >= 0; i-- {
		dp[m][i] = dp[m][i+1] + getASCII(string(s2[i]))
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = min(dp[i+1][j]+getASCII(string(s1[i])), dp[i][j+1]+getASCII(string(s2[j])))
			}
		}
	}
	return dp[0][0]
	// 根据官方思路，应该从后面往前面递推
	// 同时，初始化，后面的数据
	// // dp[i][j] 表示 s1 => i, s2 => j 之间的 ascii 字符的最小和
	// for i := 1; i <= m; i++ {
	// 	for j := 1; j <= n; j++ {
	// 		if s1[i-1] == s2[j-1] {
	// 			dp[i][j] = dp[i-1][j-1]
	// 		} else {
	// 			now := math.MaxInt32
	// 			now = min(dp[i-1][j]+getASCII(string(s1[i-1])), dp[i][j-1]+getASCII(string(s2[j-1])))
	// 			now = min(dp[i][j], now)
	// 			dp[i][j] = now
	// 		}
	// 	}
	// }
	// return dp[m][n]
}
func min3(a, b, c int) int {
	return min(a, min(b, c))
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func getASCII(str string) int {
	sum := 0
	for i := 0; i < len(str); i++ {
		sum += int(str[i])
	}
	return sum
}

// @lc code=end
// class Solution {
//     public int minimumDeleteSum(String s1, String s2) {
//         int[][] dp = new int[s1.length() + 1][s2.length() + 1];

//         for (int i = s1.length() - 1; i >= 0; i--) {
//             dp[i][s2.length()] = dp[i+1][s2.length()] + s1.codePointAt(i);
//         }
//         for (int j = s2.length() - 1; j >= 0; j--) {
//             dp[s1.length()][j] = dp[s1.length()][j+1] + s2.codePointAt(j);
//         }
//         for (int i = s1.length() - 1; i >= 0; i--) {
//             for (int j = s2.length() - 1; j >= 0; j--) {
//                 if (s1.charAt(i) == s2.charAt(j)) {
//                     dp[i][j] = dp[i+1][j+1];
//                 } else {
//                     dp[i][j] = Math.min(dp[i+1][j] + s1.codePointAt(i),
//                                         dp[i][j+1] + s2.codePointAt(j));
//                 }
//             }
//         }
//         return dp[0][0];
//     }
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/minimum-ascii-delete-sum-for-two-strings/solution/liang-ge-zi-fu-chuan-de-zui-xiao-asciishan-chu-he-/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。