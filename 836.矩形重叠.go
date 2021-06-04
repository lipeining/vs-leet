/*
 * @lc app=leetcode.cn id=836 lang=golang
 *
 * [836] 矩形重叠
 */

// @lc code=start
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	min := func(a, b int)int{
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int)int{
		if a > b {
			return a
		}
		return b
	}
	return (min(rec1[2], rec2[2]) > max(rec1[0], rec2[0]) &&
	min(rec1[3], rec2[3]) > max(rec1[1], rec2[1]))
}
// @lc code=end

