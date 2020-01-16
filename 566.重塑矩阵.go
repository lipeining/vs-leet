/*
 * @lc app=leetcode.cn id=566 lang=golang
 *
 * [566] 重塑矩阵
 */

// @lc code=start
func matrixReshape(nums [][]int, r int, c int) [][]int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	m := len(nums[0])
	if m == 0 {
		return nums
	}
	if n*m < r*c {
		return nums
	}
	res := make([][]int, 0)
	x:=0
	y:=0
	for i:=0; i<r;i++ {
		row := make([]int, 0)
		for j:=0;j<c;j++ {
			if (y==m) {
				x++
				y=0
			}
			tmp:=nums[x][y]
			y++
			row = append(row, tmp)
		}
		res = append(res, row)
	}
	return res
}
// @lc code=end

