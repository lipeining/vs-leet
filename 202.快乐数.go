/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	m := make(map[int]bool)
	m[n] = true
	getNext := func(num int) int {
		sum := 0
		for num != 0 {
			mod := num % 10
			sum += mod * mod
			num /= 10
		}
		return sum
	}
	for n != 1 {
		n = getNext(n)
		// fmt.Println("n", n)
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = true
	}
	return true
}

// @lc code=end

