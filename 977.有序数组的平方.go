/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(A []int) []int {
	ans := make([]int, 0)
	for i:=0;i<len(A);i++ {
		ans = append(ans, A[i]*A[i])
	}
	sort.Ints(ans)
	return ans
}
// @lc code=end

