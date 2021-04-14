/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l := 1
	r := n
	for l < r {
		m := l + (r-l)/2
		g := guess(m)
		if g < 0 {
			r = m - 1
		} else if g > 0 {
			l = m + 1
		} else {
			return m
		}
	}
	return l
}

// @lc code=end

