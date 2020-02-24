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
	// 	我们可以对数组 P 进行一次线性扫描，
	// 当扫描到 P[j] 时，我们需要得到的是满足
	//  P[j] = P[i] + S 且 i < j 的 i 的数目，
	// 使用计数器（map 或 dict）可以满足要求。

	// 作者：LeetCode
	// 链接：https://leetcode-cn.com/problems/binary-subarrays-with-sum/solution/he-xiang-tong-de-er-yuan-zi-shu-zu-by-leetcode/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	for _, n := range P {
		ans += m[n]
		m[n+S]++
	}
	return ans
}

// @lc code=end
