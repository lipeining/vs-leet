/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	// length := len(nums)
	// if length == 0 {
	// 	return 0
	// }
	// up, down := 1, 1
	// for i := 1; i < length; i++ {
	// 	if nums[i] > nums[i-1] {
	// 		up = down + 1
	// 	} else if nums[i] < nums[i-1] {
	// 		down = up + 1
	// 	}
	// }
	// if up > down {
	// 	return up
	// }
	// return down

	length := len(nums)
	if length == 0 {
		return 0
	}
	up := make([]int, length)
	down := make([]int, length)
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				up[i] = max(up[i], down[j]+1)
			} else if nums[i] < nums[j] {
				down[i] = max(down[i], up[j]+1)
			}
		}
	}
	return 1 + max(up[length-1], down[length-1])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
// public int wiggleMaxLength(int[] nums) {
// 	if (nums.length < 2)
// 		return nums.length;
// 	int[] up = new int[nums.length];
// 	int[] down = new int[nums.length];
// 	for (int i = 1; i < nums.length; i++) {
// 		for(int j = 0; j < i; j++) {
// 			if (nums[i] > nums[j]) {
// 				up[i] = Math.max(up[i],down[j] + 1);
// 			} else if (nums[i] < nums[j]) {
// 				down[i] = Math.max(down[i],up[j] + 1);
// 			}
// 		}
// 	}
// 	return 1 + Math.max(down[nums.length - 1], up[nums.length - 1]);
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/wiggle-subsequence/solution/bai-dong-xu-lie-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。