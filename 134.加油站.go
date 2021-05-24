/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if gas[i] < cost[i] {
			continue
		} else if helper(gas, cost, i) {
			return i
		}
	}
	return -1
}
func helper(gas []int, cost []int, start int) bool {
	sum := 0
	for i := start; i < len(gas)+start; i++ {
		mod := i % len(gas)
		sum = sum + gas[mod] - cost[mod]
		if sum < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

