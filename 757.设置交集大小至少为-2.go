/*
 * @lc app=leetcode.cn id=757 lang=golang
 *
 * [757] 设置交集大小至少为2
 */

// @lc code=start
func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})
	n := len(intervals)
	todo := make([]int, n)
	for i := 0; i < n; i++ {
		todo[i] = 2
	}
	t := n
	ans := 0
	for t >= 1 {
		t--
		s := intervals[t][0]
		m := todo[t]
		for p := s; p < s+m; p++ {
			for i := 0; i <= t; i++ {
				if todo[i] > 0 && p <= intervals[i][1] {
					todo[i]--
				}
			}
			ans++
		}
	}
	return ans
}

// @lc code=end

