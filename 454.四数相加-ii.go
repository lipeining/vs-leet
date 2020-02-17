/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(A []int, B []int, C []int, D []int) int {
	twoSum := make(map[int]int)
	for i:=0;i<len(A);i++ {
		for j:=0;j<len(B);j++ {
			sum := A[i]+B[j]
			twoSum[sum]++
		}
	}
	ans := 0
	for i:=0;i<len(C);i++ {
		for j:=0;j<len(D);j++ {
			sum := -(C[i]+D[j])
			if v,ok := twoSum[sum]; ok {
				ans += v
			}
		}
	}
	return ans
}
// @lc code=end

