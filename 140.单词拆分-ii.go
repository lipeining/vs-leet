/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	length := len(s)
	if length == 0 {
		tmp := make([]string, 0)
		return tmp
	}
	if len(wordDict) == 0 {
		tmp := make([]string, 0)
		return tmp
	}
	wordSet := make(map[string]bool)
	for k := 0; k < len(wordDict); k++ {
		wordSet[wordDict[k]] = true
	}
	dp := make([][]string, length+1)
	initial := make([]string, 0)
	initial = append(initial, "")
	dp[0] = initial
	for i := 1; i <= length; i++ {
		list := make([]string, 0)
		for j := 0; j < i; j++ {
			_, ok := wordSet[s[j:i]]
			if len(dp[j]) > 0 && ok {
				for _, l := range dp[j] {
					str := l
					if l != "" {
						str += " "
					}
					str += s[j:i]
					list = append(list, str)
				}
			}
		}
		dp[i] = list
	}
	return dp[length]
}

// @lc code=end
// public List<String> wordBreak(String s, Set<String> wordDict) {
// 	LinkedList<String>[] dp = new LinkedList[s.length() + 1];
// 	LinkedList<String> initial = new LinkedList<>();
// 	initial.add("");
// 	dp[0] = initial;
// 	for (int i = 1; i <= s.length(); i++) {
// 		LinkedList<String> list = new LinkedList<>();
// 		for (int j = 0; j < i; j++) {
// 			if (dp[j].size() > 0 && wordDict.contains(s.substring(j, i))) {
// 				for (String l : dp[j]) {
// 					list.add(l + (l.equals("") ? "" : " ") + s.substring(j, i));
// 				}
// 			}
// 		}
// 		dp[i] = list;
// 	}
// 	return dp[s.length()];
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/word-break-ii/solution/dan-ci-chai-fen-ii-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
