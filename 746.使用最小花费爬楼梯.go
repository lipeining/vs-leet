/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	sum := make([]int, length)
	sum[0] = cost[0]
	sum[1] = cost[1]
	for i:=2;i<length;i++ {
		min:= sum[i-1]
		if sum[i-1] > sum[i-2] {
			min=sum[i-2]	
		}
		sum[i]=cost[i] + min
	}
	if sum[length-1]  > sum[length-2] {
		return sum[length-2]
	}
	return sum[length-1]
}
// @lc code=end

