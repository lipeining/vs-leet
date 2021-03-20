/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, num := range nums1 {
		m[num]++
	}
	s := make(map[int]int)
	for _, num := range nums2 {
		s[num]++
	}
	if len(s) > len(m) {
		s, m = m, s
	}
	ans := make([]int, 0)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for k, v := range s {
		if v2, ok := m[k]; ok {
			for i := 0; i < min(v, v2); i++ {
				ans = append(ans, k)
			}
		}
	}
	return ans
}

// @lc code=end

