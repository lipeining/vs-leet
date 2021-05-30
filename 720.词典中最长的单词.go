/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
 */

// @lc code=start
func longestWord(words []string) string {
	smap := make(map[string]bool)
	ans := ""
	// sort.Strings(words)
	// fmt.Println(words)
	for _, word := range words {
		smap[word] = true
	}
	for _, word := range words {
		length := len(word)
		if length > len(ans) || (length == len(ans) && word < ans) {
			flag := true
			for k := 1; k < length; k++ {
				if _, ok := smap[word[0:k]]; !ok {
					flag = false
					break
				}
			}
			if flag {
				ans = word
			}
		}
	}
	return ans
}

// @lc code=end

