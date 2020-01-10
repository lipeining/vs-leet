/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */
import "fmt"
// @lc code=start
func countAndSay(n int) string {
	start := "1"
	for i:=1; i < n ; i++ {
		start = spell(start)
		// fmt.Println(start, "start")
	}
	return start
}
func spell(s string) string {
	res := ""
	length := len(s)
	for i := 0; i < length; {
		j := i
		for ; j < length - 1; j++ {
			if s[j] != s[j+1] {
				// // 得到第一个分歧点
				// l := j+1 - i
				// str := strconv.Itoa(l) + string(s[i])
				// // fmt.Println(strconv.Itoa(l), str, "-")
				// i = j + 1
				// res = res + str
				break
			}
		}
		// if j == length - 1 {
		// 	l := j+1 - i
		// 	str := strconv.Itoa(l) + string(s[i])
		// 	// fmt.Println(strconv.Itoa(l), str, "--")
		// 	i = j + 1
		// 	res = res + str
		// }
		l := j+1 - i
		str := strconv.Itoa(l) + string(s[i])
		// fmt.Println(strconv.Itoa(l), str, "--")
		i = j + 1
		res = res + str
	}
	return res;
}
// @lc code=end

