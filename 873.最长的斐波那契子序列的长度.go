/*
 * @lc app=leetcode.cn id=873 lang=golang
 *
 * [873] 最长的斐波那契子序列的长度
 */

// @lc code=start
func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	set := make(map[int]bool)
	for _, num := range arr {
		set[num] = true
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x := arr[i]
			y := arr[i] + arr[j]
			length := 2
			for set[y] {
				z := x + y
				x = y
				y = z
				length++
				ans = max(ans, length)
			}
		}
	}
	if ans >= 3 {
		return ans
	}
	return 0
}

// @lc code=end

