/*
 * @lc app=leetcode.cn id=1105 lang=golang
 *
 * [1105] 填充书架
 */
package main

// func main() {
// 	minHeightShelves([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4)
// }

// @lc code=start
func minHeightShelves(books [][]int, shelf_width int) int {
	n := len(books)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1000 * 1000
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		width, j, h := 0, i, 0
		for j > 0 {
			width += books[j-1][0]
			if width > shelf_width {
				break
			}
			h = max(h, books[j-1][1])
			dp[i] = min(dp[i], dp[j-1]+h)
			j--
		}
	}
	return dp[n]
}

// @lc code=end
