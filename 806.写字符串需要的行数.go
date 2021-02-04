/*
 * @lc app=leetcode.cn id=806 lang=golang
 *
 * [806] 写字符串需要的行数
 */

// @lc code=start
func numberOfLines(widths []int, S string) []int {
	n := len(S)
	line := 1
	length := 100
	for i := 0; i < n; i++ {
		index := S[i] - 'a'
		if length < widths[index] {
			line++
			length = 100
		}
		length -= widths[index]
	}
	return []int{line, 100 - length}
}

// @lc code=end

