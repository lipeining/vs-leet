/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	length := len(height)
	ans := 0
	left := 0
	right := length - 1
	for left < right {
		area := getArea(height, left, right)
		if area > ans {
			ans = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	// 短的进行移动
	return ans
	// length := len(height)
	// ans := 0
	// for i := 0; i < length; i++ {
	// 	for j := i + 1; j < length; j++ {
	// 		area := getArea(height, i, j)
	// 		if area > ans {
	// 			ans = area
	// 		}
	// 	}
	// }
	// return ans
}
func getArea(height []int, left int, right int) int {
	h := min(height[left], height[right])
	return h * (right - left)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
// public int maxArea(int[] height) {
// 	int maxarea = 0, l = 0, r = height.length - 1;
// 	while (l < r) {
// 		maxarea = Math.max(maxarea, Math.min(height[l], height[r]) * (r - l));
// 		if (height[l] < height[r])
// 			l++;
// 		else
// 			r--;
// 	}
// 	return maxarea;
// }

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/container-with-most-water/solution/sheng-zui-duo-shui-de-rong-qi-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。