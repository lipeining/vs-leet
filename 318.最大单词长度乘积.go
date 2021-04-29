/*
 * @lc app=leetcode.cn id=318 lang=golang
 *
 * [318] 最大单词长度乘积
 */

// @lc code=start
func maxProduct(words []string) int {
	n := len(words)
	trans := make([]int, n)
	for i, word := range words {
		num := 0
		for _, char := range word {
			num |= 1 << int(char-'a')
		}
		trans[i] = num
	}
	// fmt.Println(trans)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if trans[i]&trans[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

// @lc code=end

