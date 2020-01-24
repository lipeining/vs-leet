/*
 * @lc app=leetcode.cn id=1005 lang=golang
 *
 * [1005] K 次取反后最大化的数组和
 */

// @lc code=start
func largestSumAfterKNegations(A []int, K int) int {
	for i:=0;i<K;i++ {
		sort.Ints(A)
		A[0]=-A[0]
	}
	sum := 0
	for i:=0;i<len(A);i++ {
		sum += A[i]
	}
	return sum
}
// @lc code=end

