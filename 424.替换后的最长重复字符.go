/*
 * @lc app=leetcode.cn id=424 lang=golang
 *
 * [424] 替换后的最长重复字符
 */

// @lc code=start
func characterReplacement(s string, k int) int {
	var counter [26]int
	slow, fast := 0, 0
	result, maxCount := 0, 0
	for fast < len(s) {
		counter[s[fast]-'A']++
		maxCount = max(maxCount, counter[s[fast]-'A'])
		if fast-slow+1-maxCount > k {
			counter[s[slow]-'A']--
			slow++
		}
		result = max(result, fast-slow+1)
		fast++
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
