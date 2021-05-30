/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(T []int) []int {
	// 递减栈
	n := len(T)
	ans := make([]int, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) != 0 && T[stack[len(stack)-1]] < T[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return ans
}
func dailyTemperaturesOld(T []int) []int {
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
