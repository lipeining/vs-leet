/*
 * @lc app=leetcode.cn id=820 lang=golang
 *
 * [820] 单词的压缩编码
 */

// @lc code=start
func minimumLengthEncoding(words []string) int {
	// n := len(words)
	set := make(map[string]bool)
	for _, word := range words {
		set[word] = true
	}
	for _, word := range words {
		for start := 1; start < len(word); start++ {
			delete(set, word[start:])
		}
	}
	ans := 0
	for k := range set {
		ans += len(k) + 1
	}
	return ans
}

// @lc code=end

