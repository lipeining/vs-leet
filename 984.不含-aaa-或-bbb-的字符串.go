/*
 * @lc app=leetcode.cn id=984 lang=golang
 *
 * [984] 不含 AAA 或 BBB 的字符串
 */
package main

// func main() {
// 	strWithout3a3b(1, 2)
// 	strWithout3a3b(4, 1)
// 	strWithout3a3b(40, 20)
// }

// @lc code=start
func strWithout3a3b(a int, b int) string {
	ac, bc := 0, 0
	want := a + b
	ans := ""
	for len(ans) != want {
		if a > b && ac < 2 || a <= b && bc == 2 {
			ans += "a"
			a--
			ac++
			bc = 0
		} else {
			ans += "b"
			b--
			bc++
			ac = 0
		}
	}
	// fmt.Println("anssss", ans)
	return ans
}

// @lc code=end
