/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 */

// @lc code=start
func maxProduct(nums []int) int {

	// maxDP[i + 1] = max(maxDP[i] * A[i + 1], A[i + 1],minDP[i] * A[i + 1])
	// minDP[i + 1] = min(minDP[i] * A[i + 1], A[i + 1],maxDP[i] * A[i + 1])
	// dp[i + 1] = max(dp[i], maxDP[i + 1])

	// 作者：yang-cong-12
	// 链接：https://leetcode-cn.com/problems/maximum-product-subarray/solution/dpfang-fa-xiang-jie-by-yang-cong-12/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

	


	// ans := math.MinInt32
	// imax := 1
	// imin := 1
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] < 0 {
	// 		imax, imin = imin, imax
	// 	}
	// 	imax = max(imax*nums[i], nums[i])
	// 	imin = min(imin*nums[i], nums[i])
	// 	ans = max(ans, imax)
	// }
	// return ans
	// // 考虑 0 连续的负数情况吧
	// length := len(nums)
	// if length == 0 {
	// 	return 0
	// }
	// dp := make([]int, length)
	// dp[0] = nums[0]
	// // 考虑前缀乘积，前缀和
	// // dp[i] 表示以 i 结尾的最大值
	// for i := 1; i < length; i++ {
	// 	should := 0
	// 	for j := 0; j < i; j++ {
	// 		should = max(nums[i], nums[i])
	// 	}
	// 	dp[i] = should
	// }
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
// public int maxProduct(int[] nums) {
// 	int max = Integer.MIN_VALUE, imax = 1, imin = 1;
// 	for(int i=0; i<nums.length; i++){
// 		if(nums[i] < 0){
// 		  int tmp = imax;
// 		  imax = imin;
// 		  imin = tmp;
// 		}
// 		imax = Math.max(imax*nums[i], nums[i]);
// 		imin = Math.min(imin*nums[i], nums[i]);

// 		max = Math.max(max, imax);
// 	}
// 	return max;
// }

// 作者：guanpengchn
// 链接：https://leetcode-cn.com/problems/maximum-product-subarray/solution/hua-jie-suan-fa-152-cheng-ji-zui-da-zi-xu-lie-by-g/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。