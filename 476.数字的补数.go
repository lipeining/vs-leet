/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */

// @lc code=start
func findComplement(num int) int {
	ans := 0
	if num == 0 {
		return 1
	}
	for i := 0; i < 32 && num != 0; i++ {
		bit := num & 1
		if bit == 0 {
			ans |= 1 << i
		}
		num >>= 1
	}
	return ans
	// res := 0
	// multi := 1
	// for num != 0 {
	// 	m := num% 2
	// 	num = num / 2
	// 	if m == 0 {
	// 		res += multi
	// 	}
	// 	multi = multi * 2
	// }
	// return res
}

// @lc code=end

