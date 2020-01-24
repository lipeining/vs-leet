/*
 * @lc app=leetcode.cn id=1217 lang=golang
 *
 * [1217] 玩筹码
 */

// @lc code=start
func minCostToMoveChips(chips []int) int {
	even,odd := 0, 0
	for i:=0;i<len(chips);i++ {
		if chips[i]%2 == 1 {
			odd++
		} else {
			even++
		}
	}
	if odd > even {
		return even
	}
	return odd
}
// @lc code=end

