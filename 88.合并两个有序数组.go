/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	// 从尾部开始排序合并即可
	end := m + n
	for m > 0 || n > 0 {
		if m == 0 {
			nums1[end-1] = nums2[n-1]
			end--
			n--
			continue
		}
		if n == 0 {
			nums1[end-1] = nums1[m-1]
			end--
			m--
			continue
		}
		if nums1[m-1] >= nums2[n-1] {
			nums1[end-1] = nums1[m-1]
			end--
			m--
		} else {
			nums1[end-1] = nums2[n-1]
			end--
			n--
		}
	}
}
// @lc code=end

