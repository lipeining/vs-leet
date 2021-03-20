/*
 * @lc app=leetcode.cn id=532 lang=golang
 *
 * [532] 数组中的 k-diff 数对
 */

// @lc code=start
func findPairs(nums []int, k int) int {
	// 存储小的值
	if k < 0 {
		return 0
	}
	seen, diff := make(map[int]bool), make(map[int]bool)
	for _, num := range nums {
		if _, ok := seen[num-k]; ok {
			diff[num-k] = true
		}
		if _, ok := seen[num+k]; ok {
			diff[num] = true
		}
		seen[num] = true
	}
	return len(diff)
}

// @lc code=end

