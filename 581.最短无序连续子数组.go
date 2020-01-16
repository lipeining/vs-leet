/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
// 	int l = nums.length, r = 0;
// 	for (int i = 0; i < nums.length - 1; i++) {
// 		for (int j = i + 1; j < nums.length; j++) {
// 			if (nums[j] < nums[i]) {
// 				r = Math.max(r, j);
// 				l = Math.min(l, i);
// 			}
// 		}
// 	}
// 	return r - l < 0 ? 0 : r - l + 1;
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/solution/zui-duan-wu-xu-lian-xu-zi-shu-zu-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	length := len(nums)
	l,r:=float64(length),float64(0)
	for i:=0;i<length-1;i++ {
		for j:=i+1;j<length;j++ {
			if nums[j] <nums[i] {
				r = math.Max(r,float64(j))
				l = math.Min(l,float64(i))
			}
		}
	}
	if r < l {
		return 0
	}
	return int(r-l)+1
}
// @lc code=end

