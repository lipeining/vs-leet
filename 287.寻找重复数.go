/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	// 可能出现 2，3，4，5 次
	set := make(map[int]bool)
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return num
		} else {
			set[num] = true
		}
	}
	return 0
}

// @lc code=end
// public int findDuplicate(int[] nums) {
// 	// Find the intersection point of the two runners.
// 	int tortoise = nums[0];
// 	int hare = nums[0];
// 	do {
// 		tortoise = nums[tortoise];
// 		hare = nums[nums[hare]];
// 	} while (tortoise != hare);

// 	// Find the "entrance" to the cycle.
// 	int ptr1 = nums[0];
// 	int ptr2 = tortoise;
// 	while (ptr1 != ptr2) {
// 		ptr1 = nums[ptr1];
// 		ptr2 = nums[ptr2];
// 	}

// 	return ptr1;
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/find-the-duplicate-number/solution/xun-zhao-zhong-fu-shu-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。