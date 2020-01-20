/*
 * @lc app=leetcode.cn id=961 lang=golang
 *
 * [961] 重复 N 次的元素
 */
import "sort"
// @lc code=start
func repeatedNTimes(A []int) int {
	for k:=1;k<=3;k++ {
		for i:=0; i<len(A)-k;i++ {
			if A[i] == A[i+k] {
				return A[i]
			}
		}
	}
	return 0
}
// @lc code=end

