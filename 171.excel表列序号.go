/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */

// @lc code=start
func titleToNumber(s string) int {
	res := 0
	str := " ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length := len(s)
	multi := 1
	for i:= length-1; i>=0; i-- {
		index := strings.Index(str, string(s[i]))
		if index == -1 {
			// should not happen
		} else {
			res = res + multi * index
			multi = multi * 26
		}
	}
	return res
}
// @lc code=end

