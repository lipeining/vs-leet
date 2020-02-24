/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] 乘积小于K的子数组
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {

	if k <= 1 {
		return 0
	}
	prod := 1
	ans := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		prod *= nums[right]
		for prod >= k {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
	// length := len(nums)
	// prefix := make([]int, length+1)
	// prefix[0] = 1
	// for i := 1; i <= length; i++ {
	// 	prefix[i] = prefix[i-1] * nums[i-1]
	// }
	// ans := 0
	// for i := 0; i < length; i++ {
	// 	for j := i + 1; j <= length; j++ {
	// 		multi := prefix[j] / prefix[i]
	// 		if multi < k {
	// 			ans++
	// 		}
	// 	}
	// }
	// return ans
}

// @lc code=end
// public int numSubarrayProductLessThanK(int[] nums, int k) {
// 	if (k <= 1) return 0;
// 	int prod = 1, ans = 0, left = 0;
// 	for (int right = 0; right < nums.length; right++) {
// 		prod *= nums[right];
// 		while (prod >= k) prod /= nums[left++];
// 		ans += right - left + 1;
// 	}
// 	return ans;
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/subarray-product-less-than-k/solution/cheng-ji-xiao-yu-kde-zi-shu-zu-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。