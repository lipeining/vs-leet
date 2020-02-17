/*
 * @lc app=leetcode.cn id=930 lang=golang
 *
 * [930] 和相同的二元子数组
 */

// @lc code=start
func numSubarraysWithSum(A []int, S int) int {
	length := len(A)
	P := make([]int, length+1)
	for i := 0; i < length; i++ {
		P[i+1] = P[i] + A[i]
	}
	ans := 0
	m := make(map[int]int)
	for _, n := range P {
		ans += m[n]
		m[n+S]++
	}
	return ans
}

// @lc code=end
