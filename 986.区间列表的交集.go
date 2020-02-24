/*
 * @lc app=leetcode.cn id=986 lang=golang
 *
 * [986] 区间列表的交集
 */

// @lc code=start
func intervalIntersection(A [][]int, B [][]int) [][]int {
	ans := make([][]int, 0)
	a := 0
	b := 0
	lena := len(A)
	lenb := len(B)
	for a < lena && b < lenb {
		al, ar := A[a][0], A[a][1]
		bl, br := B[b][0], B[b][1]
		leftMax := max(al, bl)
		rightMin := min(ar, br)
		if leftMax <= rightMin {
			ans = append(ans, []int{leftMax, rightMin})
		}
		if ar < br {
			a++
		} else {
			b++
		}
	}
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
