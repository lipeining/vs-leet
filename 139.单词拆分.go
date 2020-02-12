import "fmt"

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {

	//     public boolean wordBreak(String s, List<String> wordDict) {
	//         Set<String> wordDictSet=new HashSet(wordDict);
	//         boolean[] dp = new boolean[s.length() + 1];
	//         dp[0] = true;
	//         for (int i = 1; i <= s.length(); i++) {
	//             for (int j = 0; j < i; j++) {
	//                 if (dp[j] && wordDictSet.contains(s.substring(j, i))) {
	//                     dp[i] = true;
	//                     break;
	//                 }
	//             }
	//         }
	//         return dp[s.length()];
	//     }

	// 作者：LeetCode
	// 链接：https://leetcode-cn.com/problems/word-break/solution/dan-ci-chai-fen-by-leetcode/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

	// // 根据官方解答优化代码
	// length := len(s)
	// if length == 0 {
	// 	return false
	// }
	// if len(wordDict) == 0 {
	// 	return false
	// }
	// dp := make([]bool, length+1)
	// wordSet := make(map[string]bool)
	// for k := 0; k < len(wordDict); k++ {
	// 	wordSet[wordDict[k]] = true
	// }
	// dp[0] = true
	// for i := 1; i <= length; i++ {
	// 	for j := 0; j < i; j++ {
	// 		_, ok := wordSet[s[j:i]]
	// 		if dp[j] && ok {
	// 			dp[i] = true
	// 			break
	// 		}
	// 	}
	// }
	// return dp[length]

	// 为什么 官方比较快?
	// s[j:i]
	// if (i-wL >= 0 && dp[i-wL]) || i-wL < 0 {
	// 提前知道 dp[i-wL] 是有效的在进行计算

	// 背包问题
	// wordDict 放在里层循环
	length := len(s)
	if length == 0 {
		return false
	}
	if len(wordDict) == 0 {
		return false
	}
	dp := make([]bool, length)
	for i := 0; i < length; i++ {
		for j := 0; j < len(wordDict); j++ {
			word := wordDict[j]
			wL := len(word)
			if i >= wL-1 {
				if (i-wL >= 0 && dp[i-wL]) || i-wL < 0 {
					if helper(s[i-wL+1:i+1], word) {
						if i-wL >= 0 {
							dp[i] = dp[i] || dp[i-wL]
						} else {
							dp[i] = true
						}
					}
				}
			}
			// fmt.Println(i, j, word, dp)
		}
	}
	// fmt.Println(dp)
	return dp[length-1]

	// // 背包问题
	// // wordDict 放在里层循环
	// length := len(s)
	// if length == 0 {
	// 	return false
	// }
	// if len(wordDict) == 0 {
	// 	return false
	// }
	// dp := make([]bool, length)
	// for i := 0; i < length; i++ {
	// 	for j := 0; j < len(wordDict); j++ {
	// 		word := wordDict[j]
	// 		wL := len(word)
	// 		if i >= wL-1 {
	// 			if helper(s[i-wL+1:i+1], word) {
	// 				if i-wL >= 0 {
	// 					dp[i] = dp[i] || dp[i-wL]
	// 				} else {
	// 					dp[i] = true
	// 				}
	// 			}
	// 		}
	// 		fmt.Println(i, j, word, dp)
	// 	}
	// }
	// fmt.Println(dp)
	// return dp[length-1]
}
func helper(a, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// @lc code=end
