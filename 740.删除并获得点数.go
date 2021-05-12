/*
 * @lc app=leetcode.cn id=740 lang=golang
 *
 * [740] 删除并获得点数
 */

// @lc code=start
func deleteAndEarn(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxVal := 0
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
		maxVal = max(maxVal, num)
	}
	f := make([][]int, maxVal+1)
	for i := 0; i <= maxVal; i++ {
		f[i] = make([]int, 2)
	}
	for i := 1; i <= maxVal; i++ {
		f[i][1] = f[i-1][0] + cnt[i]*i
		f[i][0] = max(f[i-1][0], f[i-1][1])
	}
	return max(f[maxVal][0], f[maxVal][1])
}

// @lc code=end

