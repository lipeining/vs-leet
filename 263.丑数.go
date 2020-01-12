/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */
import "fmt"
// @lc code=start
func isUgly(num int) bool {
	n := num
	if n<=0 {
		return false
	}
	if n == 1 {
		return true
	}
	ugly := []int{5,3,2}
	for n != 1 {
		m := -1
		for _,d := range ugly {
			m = n % d
			if m == 0 {
				n = n / d
				break
			}
		}
		// fmt.Println(m, n)
		if m > 0 {
			return false
		}
	}
	return true
}
// @lc code=end

