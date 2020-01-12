/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 */

// @lc code=start
func convertToBase7(num int) string {
	flag := false
	if num < 0 {
		flag = true
		num = -num
	}
	if num == 0 {
		return "0"
	}
	res := ""
	for num != 0 {
		m := num % 7
		num = num / 7
		res = strconv.Itoa(m) + res
	}
	if flag {
		res = "-" + res
	}
	return res
}
// @lc code=end

