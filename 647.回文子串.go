/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {
	// // for (int l = 1; l < n; ++l) {
	// // 	for (int i = 0; i < n-l; ++i) {
	// // 		int j = i + l;
	// // 		if (s[i] == s[j] && dp[i+1][j-1] == j-i-1) {
	// // 			dp[i][j] = dp[i+1][j-1] + 2;
	// // 		} else {
	// // 			dp[i][j] = 0;
	// // 		}
	// // 	}
	// // }
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 1; l < n; l++ {
		for i := 0; i < n-l; i++ {
			j := i + l
			if s[i] == s[j] && dp[i+1][j-1] == j-i-1 {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[0][n-1]
	// n := len(s)
	// // 	dp[start][end]表示以start开头,
	// // 	end结束的字符串是不是回文子串
	// // 如果 start == end,肯定是回文子串
	// // 如果 end > start,那么这个字符串是不是回文串取决于
	// // 字符串头和尾巴是否相等并且去除头尾剩余的字符串
	// // 是不是对应的回文串
	// dp := make([][]bool, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([]bool, n)
	// }
	// ans := 0
	// for end := 0; end < n; end++ {
	// 	for start := 0; start <= end; start++ {
	// 		if s[start] == s[end] && (end-start <= 1 || dp[start+1][end-1]) {
	// 			dp[start][end] = true
	// 			ans++
	// 		}
	// 	}
	// }
	// return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func max3(a, b, c int) int {
	return max(a, max(b, c))
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func min3(a, b, c int) int {
	return min(min(a, b), c)
}

// @lc code=end
// public int countSubstrings(String S) {
// 	int N = S.length(), ans = 0;
// 	for (int center = 0; center <= 2*N-1; ++center) {
// 		int left = center / 2;
// 		int right = left + center % 2;
// 		while (left >= 0 && right < N && S.charAt(left) == S.charAt(right)) {
// 			ans++;
// 			left--;
// 			right++;
// 		}
// 	}
// 	return ans;
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/palindromic-substrings/solution/hui-wen-zi-chuan-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。