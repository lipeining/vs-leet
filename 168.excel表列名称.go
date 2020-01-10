/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */
import "fmt"
// @lc code=start
func convertToTitle(n int) string {
	// 26 进制
	str := " ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length := len(str)
	res := ""
	fmt.Println(length)
	for n > 0 {
		m := n % (length-1)
		
		if m == 0 {
			m = 26
		}
		n = (n-m) / (length-1)
		res = string(str[m]) + res
	}
	return res
}
// @lc code=end

