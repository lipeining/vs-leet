/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
 */

// @lc code=start
func findNumberOfLIS(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	lengths := make([]int, length)
	counts := make([]int, length)
	for i := 0; i < length; i++ {
		counts[i] = 1
	}
	// 官方思路代码
	for j := 0; j < length; j++ {
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] {
				if lengths[i] >= lengths[j] {
					lengths[j] = lengths[i] + 1
					counts[j] = counts[i]
				} else if lengths[i]+1 == lengths[j] {
					counts[j] += counts[i]
				}
			}
		}
	}
	longest, ans := 0, 0
	for i := 0; i < length; i++ {
		longest = max(longest, lengths[i])
	}
	for i := 0; i < length; i++ {
		if lengths[i] == longest {
			ans += counts[i]
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
// Arrays.fill(counts, 1);

// for (int j = 0; j < N; ++j) {
// 	for (int i = 0; i < j; ++i) if (nums[i] < nums[j]) {
// 		if (lengths[i] >= lengths[j]) {
// 			lengths[j] = lengths[i] + 1;
// 			counts[j] = counts[i];
// 		} else if (lengths[i] + 1 == lengths[j]) {
// 			counts[j] += counts[i];
// 		}
// 	}
// }

// int longest = 0, ans = 0;
// for (int length: lengths) {
// 	longest = Math.max(longest, length);
// }
// for (int i = 0; i < N; ++i) {
// 	if (lengths[i] == longest) {
// 		ans += counts[i];
// 	}
// }
// return ans;

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/solution/zui-chang-di-zeng-zi-xu-lie-de-ge-shu-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。