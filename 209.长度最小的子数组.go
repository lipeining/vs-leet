import "math"

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	// 二分，抄思路答案
	n := len(nums)
	ans := math.MaxInt32
	left, sum := 0, 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= s {
			ans = min(ans, i-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans != math.MaxInt32 {
		return ans
	}
	return 0
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

// int ans = INT_MAX;
// int left = 0;
// int sum = 0;
// for (int i = 0; i < n; i++) {
// 	sum += nums[i];
// 	while (sum >= s) {
// 		ans = min(ans, i + 1 - left);
// 		sum -= nums[left++];
// 	}
// }
// return (ans != INT_MAX) ? ans : 0;

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum/solution/chang-du-zui-xiao-de-zi-shu-zu-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。