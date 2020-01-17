/*
 * @lc app=leetcode.cn id=657 lang=golang
 *
 * [657] 机器人能否返回原点
 */

// @lc code=start
func judgeCircle(moves string) bool {
	x := 0
	y := 0
	for _,move := range moves {
		if move == 'U' {
			y++
		} else if move == 'D' {
			y--
		} else if move == 'L' {
			x--
		} else {
			x++
		}
	}
	return x == 0 && y==0
}
// @lc code=end

