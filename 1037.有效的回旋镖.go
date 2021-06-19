/*
 * @lc app=leetcode.cn id=1037 lang=golang
 *
 * [1037] 有效的回旋镖
 */

// @lc code=start
func isBoomerang(points [][]int) bool {
	edge := func(left, right []int) (e []int) {
		return []int{right[0] - left[0], right[1] - left[1]}
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})
	e1 := edge(points[0], points[1])
	e2 := edge(points[0], points[2])
	// 90
	if e1[0] == 0 && e2[0] == 0 {
		return false
	}
	// 0
	if e1[1] == 0 && e2[1] == 0 {
		return false
	}
	// other
	return e1[0]*e2[1] != e1[1]*e2[0]
}

// @lc code=end

