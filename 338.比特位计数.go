/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(num int) []int {
	length := num
	ans := make([]int, length+1)
	for i := 0; i <= length; i++ {
		ans[i] = popcount(i)
	}
	return ans
}
func popcount(num int) int {
	counter := 0
	for num != 0 {
		num &= num - 1
		counter++
	}
	return counter
}

// @lc code=end
