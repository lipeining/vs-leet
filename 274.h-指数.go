/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H指数
 */

// @lc code=start
func hIndex(citations []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	mx := 0
	cnt := make(map[int]int)
	for _, c := range citations {
		mx = max(mx, c)
		cnt[c]++
	}
	l := 0
	r := mx
	check := func(num int) bool {
		total := 0
		for k, v := range cnt {
			if k >= num {
				total += v
			}
		}
		return total >= num
	}
	for l < r {
		mid := (l + r + 1) >> 1
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

// @lc code=end

