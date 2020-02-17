/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(T []int) []int {
	length := len(T)
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		j := i + 1
		for j < length {
			if T[j] > T[i] {
				break
			}
			j++
		}
		if j != length {
			ans[i] = j - i
		}
	}
	return ans
}

// @lc code=end
