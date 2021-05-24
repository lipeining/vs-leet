/*
 * @lc app=leetcode.cn id=502 lang=golang
 *
 * [502] IPO
 */

// @lc code=start
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	for i := 0; i < k && i < n; i++ {
		id := -1
		for j := 0; j < n; j++ {
			if w >= capital[j] {
				if id == -1 || profits[id] < profits[j] {
					id = j
				}
			}
		}
		if id == -1 {
			break
		}
		w += profits[id]
		capital[id] = math.MaxInt32
	}
	return w
}

// @lc code=end

