/*
 * @lc app=leetcode.cn id=967 lang=golang
 *
 * [967] 连续差相同的数字
 */

// @lc code=start
func numsSameConsecDiff(n int, k int) []int {
	ans := make([]int, 0)
	// abs := func(a int) int {
	// 	if a < 0 {
	// 		return -a
	// 	}
	// 	return a
	// }
	var dfs func(index, prev, num int)
	dfs = func(index, prev, num int) {
		if index == n {
			ans = append(ans, num)
			return
		}
		s := prev - k
		if s >= 0 && s <= 9 {
			dfs(index+1, s, num*10+s)
		}
		p := prev + k
		if s != p && p >= 0 && p <= 9 {
			dfs(index+1, p, num*10+p)
		}
	}
	for i := 1; i <= 9; i++ {
		dfs(1, i, i)
	}
	return ans
}

// @lc code=end

