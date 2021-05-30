/*
 * @lc app=leetcode.cn id=1029 lang=golang
 *
 * [1029] 两地调度
 */

// @lc code=start
func twoCitySchedCost(costs [][]int) int {
	divs := make([]int, 0)
	ans := 0
	for i := 0; i < len(costs); i++ {
		ans += costs[i][0]
		divs = append(divs, costs[i][1]-costs[i][0])
	}
	sort.Ints(divs)
	for i := 0; i < len(divs)/2; i++ {
		ans += divs[i]
	}
	return ans
}

// @lc code=end

