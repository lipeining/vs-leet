/*
 * @lc app=leetcode.cn id=525 lang=golang
 *
 * [525] 连续数组
 */

// @lc code=start
func findMaxLength(nums []int) int {
	m := make(map[int]int)
	ans := 0
	m[0] = -1
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count--
		} else {
			count++
		}
		if old, ok := m[count]; ok {
			want := i - old
			if want > ans {
				ans = want
			}
		} else {
			m[count] = i
		}
	}
	return ans
}

// func findMaxLength(nums []int) int {
// 	m := make(map[int]int)
// 	ans := 0
// 	for i := 0; i < len(nums); i++ {
// 		m[nums[i]]++
// 		want := check(m)
// 		if want > ans {
// 			ans = want
// 		}
// 	}
// 	return ans * 2
// }
// func check(m map[int]int) int {
// 	zero, zok := m[0]
// 	one, ook := m[1]
// 	if !zok || !ook {
// 		return 0
// 	}
// 	return min(zero, one)
// }
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end
