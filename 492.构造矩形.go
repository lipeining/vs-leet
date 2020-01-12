/*
 * @lc app=leetcode.cn id=492 lang=golang
 *
 * [492] 构造矩形
 */

// @lc code=start
func constructRectangle(area int) []int {
	max := int(math.Sqrt(float64(area)))
	for i:= max; i>=1;i-- {
		n := area % i
		if n == 0 {
			j := area / i
			if j < i {
				return []int{i, j}
			} else {
				return []int{j, i}
			}
			
		}
	}
	return []int{area, 1}
}
// @lc code=end

