/*
 * @lc app=leetcode.cn id=539 lang=golang
 *
 * [539] 最小时间差
 */

// @lc code=start
func findMinDifference(timePoints []string) int {
	minutes := func(str string) int {
		arr := strings.Split(str, ":")
		h, _ := strconv.Atoi(arr[0])
		m, _ := strconv.Atoi(arr[1])
		return h*60 + m
	}
	times := make([]int, len(timePoints))
	for i, t := range timePoints {
		times[i] = minutes(t)
	}
	sort.Ints(times)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	ans := 60 * 24
	for i := 0; i < len(times); i++ {
		if i == len(times)-1 {
			ans = min(ans, times[0]+60*24-times[i])
		} else {
			ans = min(ans, times[i+1]-times[i])
		}
	}
	return ans
}

// @lc code=end

