/*
 * @lc app=leetcode.cn id=969 lang=golang
 *
 * [969] 煎饼排序
 */

// @lc code=start
func pancakeSort(A []int) []int {
	// 设有n个煎饼，将最大的煎饼翻转到arr的最后，再对前n-1个煎饼进行相同的翻转操作。（缩小问题规模）

	// 每次将最大的煎饼翻转到最后，需要两步翻转：

	// 找到最大值的索引idx，并将idx之前的煎饼翻转；
	// 将arr整体翻转；（除了最后已经翻转完的煎饼）。
}

// @lc code=end

