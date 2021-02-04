/*
 * @lc app=leetcode.cn id=1147 lang=golang
 *
 * [1147] 段式回文
 */

// @lc code=start
func longestDecomposition(text string) int {
	// 局部贪心即可
	// 他人答案
	n := len(text)
	i, j := 0, n-1
	left, right, ans := "", "", 0
	for i < j {
		left += string(text[i])
		right = string(text[j]) + right
		if left == right {
			left = ""
			right = ""
			ans += 2
		}
		i++
		j--
	}
	if n%2 == 1 || left != "" {
		ans++
	}
	return ans
}

// @lc code=end

