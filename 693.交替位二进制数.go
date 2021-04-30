/*
 * @lc app=leetcode.cn id=693 lang=golang
 *
 * [693] 交替位二进制数
 */

// @lc code=start
func hasAlternatingBits(n int) bool {
	prev := -1
	for n != 0 {
		now := n & 1
		if prev == -1 {
			prev = now
		} else {
			if prev == now {
				return false
			}
			prev = now
		}
		n >>= 1
	}
	return true
	// before := 0 
	// if n % 2 ==0 {
	// 	before = 1
	// }
	// for n != 0 {
	// 	m := n % 2
	// 	n = n / 2
	// 	if before == m {
	// 		return false
	// 	} else {
	// 		before = m
	// 	}
	// }
	// return true
}
// @lc code=end

