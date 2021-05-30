/*
 * @lc app=leetcode.cn id=762 lang=golang
 *
 * [762] 二进制表示中质数个计算置位
 */

// @lc code=start
func countPrimeSetBits(left int, right int) int {
	count := func(num int) int {
		ret := 0
		for num != 0 {
			if num&1 != 0 {
				ret++
			}
			num >>= 1
		}
		return ret
	}
	ans := 0
	for index := left; index <= right; index++ {
		cnt := count(index)
		switch cnt {
		case 2, 3, 5, 7, 11, 13, 17, 19:
			ans++
		default:
			continue
		}
	}
	return ans
}

// @lc code=end

