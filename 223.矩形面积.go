/*
 * @lc app=leetcode.cn id=223 lang=golang
 *
 * [223] 矩形面积
 */

// @lc code=start
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	if A > E {
		return computeArea(E, F, G, H, A, B, C, D)
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	max := func(a, b int)int{
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int)int{
		if a < b {
			return a
		}
		return b
	}
	area := func(a, b int, c, d int) int {
		return abs(a-c) * abs(b-d)
	}
	ans := area(A, B, C, D) + area(E, F, G, H)
	if B >= H || D <= F || C <= E {
		return ans
	}
	left := max(A, E)
	down := max(B, F)
	right := min(C, G)
	up := min(D, H)
	// 如何判断是否相交
	mix := area(left, down, right, up)
	return ans - mix
}

// @lc code=end

