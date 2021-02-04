/*
 * @lc app=leetcode.cn id=964 lang=golang
 *
 * [964] 表示数字的最少运算符
 */

// @lc code=start
func leastOpsExpressTarget(x int, target int) int {
	cost := func(x int) int {
		if x > 0 {
			return x
		}
		return 2
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	memo := make(map[string]int)
	var dp func(i, t int) int
	dp = func(i, t int) int {
		code := "" + strconv.Itoa(i) + "#" + strconv.Itoa(t)
		if memo[code] != 0 {
			return memo[code]
		}
		ans := 0
		if t == 0 {
			ans = 0
		} else if t == 1 {
			ans = cost(i)
		} else if i >= 39 {
			ans = t + 1
		} else {
			nt := t / x
			r := t % x
			ans = min(r*cost(i)+dp(i+1, nt), (x-r)*cost(i)+dp(i+1, nt+1))
		}
		memo[code] = ans
		return ans
	}
	return dp(0, target) - 1
}

// @lc code=end

