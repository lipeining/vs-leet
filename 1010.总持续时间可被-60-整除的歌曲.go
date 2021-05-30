/*
 * @lc app=leetcode.cn id=1010 lang=golang
 *
 * [1010] 总持续时间可被 60 整除的歌曲
 */

// @lc code=start
func numPairsDivisibleBy60(time []int) int {
	second := make([]int, 60)
	for _, t := range time {
		second[t%60]++
	}
	ans := 0
	ans += second[0] * (second[0] - 1) / 2
	ans += second[30] * (second[30] - 1) / 2
	l, r := 1, 59
	for l < r {
		ans += second[l] * second[r]
		l++
		r--
	}
	return ans
}

// @lc code=end

