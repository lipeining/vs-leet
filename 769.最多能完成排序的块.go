/*
 * @lc app=leetcode.cn id=769 lang=golang
 *
 * [769] 最多能完成排序的块
 */

// @lc code=start
func maxChunksToSorted(arr []int) int {
	n := len(arr)
	ans := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	cur := -1
	for i := 0; i < n; i++ {
		cur = max(cur, arr[i])
		if cur == i {
			ans++
		}
	}
	return ans
}

// @lc code=end

