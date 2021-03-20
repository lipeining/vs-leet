/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for _, num := range nums1 {
		m[num] = true
	}
	ans := make([]int, 0)
	s := make(map[int]bool)
	for _, num := range nums2 {
		if m[num] {
			s[num]  = true
		}
	}
	for k,_ := range s {
		ans = append(ans, k)
	}
	return ans
}
// @lc code=end

