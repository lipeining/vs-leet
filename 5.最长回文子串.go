/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	N := len(s)
	if N < 2 {
		return s
	}
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		dp[i][i] = 1
		// if i < N-1 && s[i] == s[i+1] {
		// 	dp[i][i+1] = 1
		// }
	}
	// 如果想优化 dp 数组，需要画图，确定一维的范围和计算顺序（左右哪个先）

	// 初始化一字母和二字母的回文，然后找到所有三字母回文，并依此类推
	// p[i][j] = p[i+1][j-1] && s[i] == s[j]
	// p[i][j] = true
	// p[i][i+1] = true if s[i] == s[i+1]
	start, maxLen := 0, 1
	for j := 1; j < N; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] == 1 {
				len := j - i + 1
				if len > maxLen {
					start = i
					maxLen = len
				}
			}
		}
	}
	// fmt.Println(dp, start, maxLen)
	return s[start : start+maxLen]
}

// @lc code=end

// func longestPalindrome(s string) string {
// 	if len(s) == 0 {
// 		return ""
// 	}
// 	start, end := 0, 0
// 	for i := 0; i < len(s); i++ {
// 		len1 := expandAroundCenter(s, i, i)
// 		len2 := expandAroundCenter(s, i, i+1)
// 		len := len1
// 		if len < len2 {
// 			len = len2
// 		}
// 		if len > (end - start) {
// 			// 这里的坐标需要画图理解
// 			start = i - (len-1)/2
// 			end = i + len/2
// 		}
// 	}
// 	return s[start : end+1]
// }
// func expandAroundCenter(str string, left int, right int) int {
// 	l := left
// 	r := right
// 	for l >= 0 && r < len(str) && str[l] == str[r] {
// 		l--
// 		r++
// 	}
// 	return r - l - 1
// }
// 中心扩展
// public String longestPalindrome(String s) {
//     if (s == null || s.length() < 1) return "";
//     int start = 0, end = 0;
//     for (int i = 0; i < s.length(); i++) {
//         int len1 = expandAroundCenter(s, i, i);
//         int len2 = expandAroundCenter(s, i, i + 1);
//         int len = Math.max(len1, len2);
//         if (len > end - start) {
//             start = i - (len - 1) / 2;
//             end = i + len / 2;
//         }
//     }
//     return s.substring(start, end + 1);
// }

// private int expandAroundCenter(String s, int left, int right) {
//     int L = left, R = right;
//     while (L >= 0 && R < s.length() && s.charAt(L) == s.charAt(R)) {
//         L--;
//         R++;
//     }
//     return R - L - 1;
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。