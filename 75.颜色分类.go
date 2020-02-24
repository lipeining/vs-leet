/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	// m := make(map[int]int)
	// for _, num := range nums {
	// 	m[num]++
	// }
	// i := 0
	// list := []int{0, 1, 2}
	// for _, k := range list {
	// 	v, _ := m[k]
	// 	for v != 0 {
	// 		nums[i] = k
	// 		i++
	// 		v--
	// 	}
	// }
	p0, curr := 0, 0
	length := len(nums)
	p2 := length - 1
	tmp := 0
	for curr <= p2 {
		if nums[curr] == 0 {
			tmp = nums[p0]
			nums[p0] = nums[curr]
			nums[curr] = tmp
			p0++
			curr++
		} else if nums[curr] == 2 {
			tmp = nums[curr]
			nums[curr] = nums[p2]
			nums[p2] = tmp
			p2--
		} else {
			curr++
		}
	}
}

// @lc code=end
// class Solution {
// 	/*
// 	荷兰三色旗问题解
// 	*/
// 	public void sortColors(int[] nums) {
// 	  // 对于所有 idx < i : nums[idx < i] = 0
// 	  // j是当前考虑元素的下标
// 	  int p0 = 0, curr = 0;
// 	  // 对于所有 idx > k : nums[idx > k] = 2
// 	  int p2 = nums.length - 1;

// 	  int tmp;
// 	  while (curr <= p2) {
// 		if (nums[curr] == 0) {
// 		  // 交换第 p0个和第curr个元素
// 		  // i++，j++
// 		  tmp = nums[p0];
// 		  nums[p0++] = nums[curr];
// 		  nums[curr++] = tmp;
// 		}
// 		else if (nums[curr] == 2) {
// 		  // 交换第k个和第curr个元素
// 		  // p2--
// 		  tmp = nums[curr];
// 		  nums[curr] = nums[p2];
// 		  nums[p2--] = tmp;
// 		}
// 		else curr++;
// 	  }
// 	}
//   }

//   作者：LeetCode
//   链接：https://leetcode-cn.com/problems/sort-colors/solution/yan-se-fen-lei-by-leetcode/
//   来源：力扣（LeetCode）
//   著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。