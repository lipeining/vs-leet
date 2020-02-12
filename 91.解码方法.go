/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {

	// int numDecodings(string s) {
	// 	if (s[0] == '0') return 0;
	// 	int pre = 1, curr = 1;//dp[-1] = dp[0] = 1
	// 	for (int i = 1; i < s.size(); i++) {
	// 		int tmp = curr;
	// 		if (s[i] == '0')
	// 			if (s[i - 1] == '1' || s[i - 1] == '2') curr = pre;
	// 			else return 0;
	// 		else if (s[i - 1] == '1' || (s[i - 1] == '2' && s[i] >= '1' && s[i] <= '6'))
	// 			curr = curr + pre;
	// 		pre = tmp;
	// 	}
	// 	return curr;
	// }

	// 作者：pris_bupt
	// 链接：https://leetcode-cn.com/problems/decode-ways/solution/c-wo-ren-wei-hen-jian-dan-zhi-guan-de-jie-fa-by-pr/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	length := len(s)
	if length == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	pre, curr := 1, 1
	for i := 1; i < length; i++ {
		tmp := curr
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				curr = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			curr = curr + pre
		}
		pre = tmp
	}
	return curr

	// // dp[i] = dp[i-1] + 1 + ... + 1 假如可以得到 dp[i-2]
	// // 以 i 结尾的字符串

	// // 特殊的用例 0 00 100
	// length := len(s)
	// if length == 0 {
	// 	return 0
	// }
	// dp := make([]int, length)
	// if s[0] == '0' {
	// 	return 0
	// }
	// dp[0] = 1
	// for i := 1; i < length; i++ {
	// 	if s[i] == '0' {
	// 		dp[i] = dp[i-1]
	// 	} else {
	// 		sum := dp[i-1]
	// 		num, ok := strconv.Atoi(s[i-1 : i+1])
	// 		if ok == nil && num <= 26 && num >= 1 {
	// 			sum++
	// 		}
	// 		dp[i] = sum
	// 		fmt.Println(sum)
	// 	}
	// }
	// fmt.Println(dp)
	// return dp[length-1]
}

// @lc code=end
