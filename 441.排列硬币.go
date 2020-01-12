/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
func arrangeCoins(n int) int {
	i := 1
    for  n > 0{
		n -= i
		i++
	}
	if n == 0 {
		return i-1
	}
	return i-2
}
// @lc code=end

