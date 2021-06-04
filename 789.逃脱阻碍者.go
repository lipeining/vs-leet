/*
 * @lc app=leetcode.cn id=789 lang=golang
 *
 * [789] 逃脱阻碍者
 */

// @lc code=start
func escapeGhosts(ghosts [][]int, target []int) bool {
	source := []int{0, 0}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	taxi := func(a, b []int) int {
		return abs(a[0]-b[0]) + abs(a[1]-b[1])
	}
	for _, g := range ghosts {
		if taxi(g, target) <= taxi(source, target) {
			return false
		}
	}
	return true
}

// @lc code=end

