/*
 * @lc app=leetcode.cn id=517 lang=golang
 *
 * [517] 超级洗衣机
 */

// @lc code=start
func findMinMoves(machines []int) int {
	total := 0
	for _, num := range machines {
		total += num
	}
	n := len(machines)
	if total%n != 0 {
		return -1
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := 0
	should := total / n
	maxSum := 0
	curSum := 0
	for i := 0; i < n; i++ {
		curSum += machines[i] - should
		maxSum = max(maxSum, abs(curSum))
		ans = max(ans, max(maxSum, machines[i]-should))
	}
	return ans
}

// @lc code=end

