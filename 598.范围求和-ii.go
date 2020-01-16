/*
 * @lc app=leetcode.cn id=598 lang=golang
 *
 * [598] 范围求和 II
 */

// @lc code=start
func maxCount(m int, n int, ops [][]int) int {
    x := float64(m)
	y := float64(n)
	for i:=0;i<len(ops);i++ {
		x = math.Min(x, float64(ops[i][0]))
		y = math.Min(y, float64(ops[i][1]))
	}
	return int(x) * int(y)
}
// @lc code=end

